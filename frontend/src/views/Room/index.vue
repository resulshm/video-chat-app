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
const pc = new RTCPeerConnection({
  iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
});
const remoteVideo: any = ref();
const localVideo: any = ref();
pc.onicecandidate = (event) => {
  event.candidate &&
    sendMessage({ candidate: event.candidate, type: "candidate" });
};

onMounted(async () => {
  const allDevices = await navigator.mediaDevices.enumerateDevices();
  const cameras = allDevices.filter((device) => device.kind == "videoinput");
  const localStream = await navigator.mediaDevices.getUserMedia({
    audio: true,
    video: {
      deviceId: cameras[1].deviceId,
    },
  });
  localStream.getTracks().forEach((track: any) => {
    pc.addTrack(track, localStream);
  });
  localVideo.value.srcObject = localStream;

  pc.ontrack = (e) => {
    remoteVideo.value.srcObject = e.streams[0];
  };

  conn.value = new WebSocket(
    `ws://127.0.0.1:8080/api/v1/join/${route.params.id}`
  );

  conn.value.onopen = () => {
    sendMessage({ joined: true });
  };

  conn.value.onerror = (error: any) => {
    console.error(error);
  };

  conn.value.onmessage = async (e: any) => {
    const message = JSON.parse(e.data);
    console.log(message);
    if (message.joined) {
      await createOffer();
    }
    if (message.type === "offer") {
      answerOffer(message.offer);
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

function handleAnswer(answer: any) {
  pc.setRemoteDescription(new RTCSessionDescription(answer));
}

async function handleIceCandidate(candidate: any) {
  await pc.addIceCandidate(new RTCIceCandidate(candidate));
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
    <video ref="remoteVideo" autoplay controls></video>
    <video ref="localVideo" autoplay controls></video>
  </div>
</template>
