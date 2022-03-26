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
let localStream: any = null;
let remoteStream: any = null;

const servers = {
  iceServers: [
    {
      urls: ["stun:stun1.l.google.com:19302", "stun:stun2.l.google.com:19302"],
    },
  ],
  iceCandidatePoolSize: 10,
};
const pc = new RTCPeerConnection(servers);
pc.onicecandidate = (event) => {
  event.candidate &&
    sendMessage({ candidate: event.candidate, type: "candidate" });
};
onMounted(() => {
  // handleMedia();
  conn.value = new WebSocket(
    `ws://127.0.0.1:8080/api/v1/join/${route.params.id}`
  );
  conn.value.onopen = () => {
    sendMessage({ join: true });
    console.log("Connected");
  };
  conn.value.onerror = (error: any) => {
    console.error(error);
  };
  conn.value.onmessage = (e: any) => {
    console.log("Got message", e.data);
    const message = JSON.parse(e.data);
    if (message.join) {
      console.log("Joined bir pituh");
      createOffer();
    }
    if (message.type === "offer") {
      answerOffer(message.offer);
    }
    if (message === "answer") {
      handleAnswer(message.answer);
    }
    if (message.type === "candidate") {
      handleIceCandidate(message.candidate);
    }
  };
});

function sendMessage(message: any) {
  conn.value.send(JSON.stringify(message));
}

function handleMedia() {
  localStream = navigator.mediaDevices.getUserMedia({
    audio: true,
    video: true,
  });

  remoteStream = new MediaStream();

  localStream
    .getTracks()
    .forEach((track: any) => pc.addTrack(track, localStream));

  pc.ontrack = (event) => {
    event.streams[0].getTracks().forEach((track: any) => {
      remoteStream.addTrack(track);
    });
  };
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
  <div></div>
</template>
