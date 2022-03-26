<script lang="ts">
export default {
  name: "RoomView",
};
</script>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
const route = useRoute();
const conn = ref();
let remoteStream: any = ref();
let localStream: any = ref();
let localVideo: any = ref();
const pc = new RTCPeerConnection({
  iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
});

pc.onicecandidate = (event) => {
  event.candidate &&
    sendMessage({ candidate: event.candidate, type: "candidate" });
};
onMounted(async () => {
  localStream.value = await navigator.mediaDevices.getUserMedia({
    audio: true,
    video: true,
  });

  localStream.value.getTracks().forEach((track: any) => {
    pc.addTrack(track, localStream.value);
    localVideo.value.srcObject = track;
  });
  pc.ontrack = (e) => {
    remoteStream.value.srcObject = e.streams[0];
    console.log("Remote video", remoteStream.value);
    console.log(pc);
  };
  conn.value = new WebSocket(
    `ws://127.0.0.1:8080/api/v1/join/${route.params.id}`
  );
  conn.value.onopen = () => {
    sendMessage({ join: true });
  };
  conn.value.onerror = (error: any) => {
    console.error(error);
  };
  conn.value.onmessage = async (e: any) => {
    const message = JSON.parse(e.data);
    if (message.join) {
      console.log("Joined");
      await createOffer();
    }
    if (message.type === "offer") {
      await answerOffer(message.offer);
    }
    if (message.type === "answer") {
      await handleAnswer(message.answer);
    }
    if (message.type === "candidate") {
      await handleIceCandidate(message.candidate);
    }
  };
});

function sendMessage(message: any) {
  conn.value.send(JSON.stringify(message));
}

async function createOffer() {
  const offer = await pc.createOffer();
  await pc.setLocalDescription(offer);
  sendMessage({ offer, type: "offer" });
}

async function handleAnswer(answer: any) {
  pc.setRemoteDescription(new RTCSessionDescription(answer));
}

async function handleIceCandidate(candidate: any) {
  pc.addIceCandidate(new RTCIceCandidate(candidate));
}

async function answerOffer(offer: any) {
  await pc.setRemoteDescription(new RTCSessionDescription(offer));
  const answer = await pc.createAnswer();
  await pc.setLocalDescription(answer);

  sendMessage({ answer, type: "answer" });
}
</script>

<template>
  <div>
    <video :src="remoteStream" autoplay controls></video>
    <video :src="localVideo" autoplay controls></video>
  </div>
</template>
