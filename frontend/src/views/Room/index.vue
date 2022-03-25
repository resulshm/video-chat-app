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
const callInf: any = {
  offerCandidates: [],
  offer: {
    sdp: null,
    type: null,
  },
  answerCandidates: [],
  answer: {
    sdp: null,
    type: null,
  },
};
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

onMounted(() => {
  conn.value = new WebSocket(
    `ws://127.0.0.1:8080/api/v1/join/${route.params.id}`
  );
  conn.value.onopen = () => {
    sendMessage({ join: true });
  };
  conn.value.onerror = (error: any) => {
    console.error(error);
  };

  conn.value.onmessage = (msg: any) => {
    console.log("Got message", msg.data);
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
  pc.onicecandidate = (event) => {
    event.candidate && callInf.offerCandidates.add(event.candidate.toJSON());
  };

  const offerDescription = await pc.createOffer();
  await pc.setLocalDescription(offerDescription);

  callInf.offer.sdp = offerDescription.sdp;
  callInf.offer.type = offerDescription.type;
  // sendMessage();
}

async function handleAnswer(answer: any) {
  if (!pc.currentRemoteDescription && answer) {
    const answerDescription = new RTCSessionDescription(answer);
    pc.setRemoteDescription(answerDescription);
  }
}

async function handleIceCandidate(data: any) {
  const candidate = new RTCIceCandidate(data);
  pc.addIceCandidate(candidate);
}

async function answerOffer() {
  pc.onicecandidate = (event) => {
    event.candidate && callInf.answerCandidates.add(event.candidate.toJSON());
  };
  const offerDescription = callInf.offer;
  await pc.setRemoteDescription(new RTCSessionDescription(offerDescription));

  const answerDescription = await pc.createAnswer();
  await pc.setLocalDescription(answerDescription);

  const answer = {
    type: answerDescription.type,
    sdp: answerDescription.sdp,
  };

  // sendMessage();
}
</script>

<template>
  <div></div>
</template>
