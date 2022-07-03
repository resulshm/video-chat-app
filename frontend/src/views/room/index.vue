<script lang="ts">
export default {
  name: "RoomView",
};
</script>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import type { Message } from "@/types/index";
const route = useRoute();
const conn = ref();
const pc = new RTCPeerConnection({
  iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
});
const remoteVideo = ref<HTMLVideoElement | null>();
const localVideo = ref<HTMLVideoElement | null>();
pc.onicecandidate = (event) => {
  event.candidate &&
    sendMessage({ candidate: event.candidate, type: "candidate" });
};

onMounted(async () => {
  const localStream = await navigator.mediaDevices.getUserMedia({
    audio: true,
    video: true,
  });
  localStream.getTracks().forEach((track: any) => {
    pc.addTrack(track, localStream);
  });
  if (localVideo.value) localVideo.value.srcObject = localStream;

  pc.ontrack = (e) => {
    if (remoteVideo.value) remoteVideo.value.srcObject = e.streams[0];
  };
  conn.value = new WebSocket(
    `${import.meta.env.VITE_WEBSOCKET_URL}/api/v1/join/${route.params.id}`
  );

  conn.value.onopen = () => {
    sendMessage({ joined: true });
  };

  conn.value.onerror = (error: any) => {
    console.error(error);
  };

  conn.value.onmessage = async (e: MessageEvent) => {
    try {
      let message: Message = await JSON.parse(e.data);
      if (message.joined) {
        createOffer();
      }
      if (message.type === "offer" && message.offer) {
        answerOffer(message.offer);
      }
      if (message.type === "answer" && message.answer) {
        handleAnswer(message.answer);
      }
      if (message.type === "candidate" && message.candidate) {
        handleIceCandidate(message.candidate);
      }
      if (message.leave) {
        handleLeave();
      }
    } catch (err) {
      console.error("Can't handle message", err);
    }
  };
});

function sendMessage(message: any) {
  if (conn.value.readyState === conn.value.OPEN) {
    conn.value.send(JSON.stringify(message));
  } else {
    console.error("No connection");
  }
}

async function createOffer() {
  try {
    const offer = await pc.createOffer();
    await pc.setLocalDescription(offer);
    sendMessage({ offer, type: "offer" });
  } catch (err) {
    console.error("Can't create offer", err);
  }
}

function handleAnswer(answer: RTCSessionDescriptionInit) {
  pc.setRemoteDescription(new RTCSessionDescription(answer));
}

async function handleIceCandidate(candidate: RTCIceCandidateInit) {
  try {
    await pc.addIceCandidate(new RTCIceCandidate(candidate));
  } catch (err) {
    console.error("Can't handle Ice candidate", err);
  }
}

async function answerOffer(offer: RTCSessionDescriptionInit) {
  try {
    await pc.setRemoteDescription(new RTCSessionDescription(offer));
    const answer = await pc.createAnswer();
    await pc.setLocalDescription(answer);
    sendMessage({ answer, type: "answer" });
  } catch (err) {
    console.error("Can't anwser");
  }
}

function handleLeave() {
  remoteVideo.value = null;
  localVideo.value = null;
  pc.close();
  pc.onicecandidate = null;
  pc.ontrack = null;
  sendMessage({ leave: true });
}
</script>

<template>
  <div class="w-full h-full flex flex-row justify-center items-center">
    <div>
      <video ref="remoteVideo" autoplay controls></video>
    </div>
    <div>
      <video ref="localVideo" autoplay controls></video>
    </div>
  </div>
</template>
