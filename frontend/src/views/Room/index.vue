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
pc.onicecandidate = handleIceCandidate;
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

  conn.value.onmessage = (e: any) => {
    console.log("Got message", e.data);
    const message = JSON.parse(e.data);
    if (message.join) {
      createOffer();
    }
    if (message.callInf.offer) {
      answerOffer(message.callInf);
    }
    if (message.callInf.answer) {
      handleAnswer(message.callInf);
    }
    if (message.icecandidate) {
      const candidate = new RTCIceCandidate(message.icecandidate);
      pc.addIceCandidate(message.icecandidate);
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
  const offer: any = {};

  const offerDescription = await pc.createOffer();
  await pc.setLocalDescription(offerDescription);

  offer.sdp = offerDescription.sdp;
  offer.type = offerDescription.type;
  sendMessage({ offer });
}

async function handleAnswer(answer: any) {
  if (!pc.currentRemoteDescription && answer) {
    const answerDescription = new RTCSessionDescription(answer);
    pc.setRemoteDescription(answerDescription);
  }
}

async function handleIceCandidate(event: any) {
  event.candidate && sendMessage({ candidate: event.candidate });
}

async function answerOffer(offer: any) {
  await pc.setRemoteDescription(new RTCSessionDescription(offer));

  const answerDescription = await pc.createAnswer();
  await pc.setLocalDescription(answerDescription);

  const answer = {
    type: answerDescription.type,
    sdp: answerDescription.sdp,
  };

  sendMessage({ answer });
}
</script>

<template>
  <div></div>
</template>
