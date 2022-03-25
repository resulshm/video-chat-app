<script lang="ts">
export default {
  name: "RoomView",
};
</script>

<script setup lang="ts">
const servers = {
  iceServers: [
    {
      urls: ["stun:stun1.l.google.com:19302", "stun:stun2.l.google.com:19302"],
    },
  ],
  iceCandidatePoolSize: 10,
};
const pc = new RTCPeerConnection(servers);
let conn = new WebSocket(`ws://localhost:8080`);
let localStream: any = null;
let remoteStream: any = null;

const callDoc: any = {
  name: "",
  offerCandidates: [],
  answerCandidates: [],
  offer: {
    sdp: null,
    type: null,
  },
};

function sendMessage() {
  // Send message to the server via Websocket
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
    event.candidate && callDoc.offerCandidates.add(event.candidate.toJSON());
  };

  const offerDescription = await pc.createOffer();
  await pc.setLocalDescription(offerDescription);

  callDoc.offer.sdp = offerDescription.sdp;
  callDoc.offer.type = offerDescription.type;
  sendMessage();
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
    event.candidate && callDoc.answerCandidates.add(event.candidate.toJSON());
  };
  const offerDescription = callDoc.offer;
  await pc.setRemoteDescription(new RTCSessionDescription(offerDescription));

  const answerDescription = await pc.createAnswer();
  await pc.setLocalDescription(answerDescription);

  const answer = {
    type: answerDescription.type,
    sdp: answerDescription.sdp,
  };

  sendMessage();
}
</script>

<template>
  <div></div>
</template>
