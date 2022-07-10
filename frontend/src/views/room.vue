<script lang="ts">
export default {
  name: "RoomView",
};
</script>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useFetch } from "@vueuse/core";
import type { Message } from "@/types";

const pc = new RTCPeerConnection({
  iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
});
pc.onicecandidate = (event) => {
  event.candidate &&
    sendMessage({ message: event.candidate, type: "candidate" });
};
pc.ontrack = (e) => {
  if (remoteVideo.value) remoteVideo.value.srcObject = e.streams[0];
};

const dialogState = ref(false);
const username = ref("");
const friendMail = ref("");
const route = useRoute();
const router = useRouter();
const conn = ref();
let localStream: null | MediaStream = null;
const remoteVideo = ref<HTMLVideoElement | null>();
const localVideo = ref<HTMLVideoElement | null>();
const messageHandler = {
  join: createOffer,
  offer: answerOffer,
  answer: handleAnswer,
  candidate: handleIceCandidate,
  leave: handleLeave,
};

onMounted(async () => {
  try {
    localStream = await navigator.mediaDevices.getUserMedia({
      audio: true,
      video: true,
    });
    localStream.getTracks().forEach((track: any) => {
      pc.addTrack(track, localStream!);
    });
  } catch (err) {
    console.error(err);
  }

  if (localVideo.value) localVideo.value.srcObject = localStream;

  conn.value = new WebSocket(
    `${import.meta.env.VITE_WEBSOCKET_URL}/api/v1/join/${route.params.id}`
  );

  conn.value.onopen = () => {
    sendMessage({ type: "join" });
  };

  conn.value.onerror = (error: any) => {
    console.error(error);
  };

  conn.value.onmessage = async (e: MessageEvent) => {
    let { type, message }: Message = await JSON.parse(e.data);
    messageHandler[type](message);
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
    sendMessage({ message: offer, type: "offer" });
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
    sendMessage({ message: answer, type: "answer" });
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
  sendMessage({ type: "leave" });
  if (localStream) {
    localStream.getTracks().forEach(function (track) {
      track.stop();
    });
  }
  router.push("/");
}

async function invite() {
  dialogState.value = false;
  const formData = new FormData();
  formData.append("from", username.value);
  formData.append("to", friendMail.value);
  formData.append("code", route.params.id as string);
  const { error } = await useFetch("/api/v1/invite", {
    method: "POST",
    body: formData,
  });
  if (error.value) {
    console.error(error);
    return;
  }
  alert("Invitation is sended");
}
</script>

<template>
  <div class="container relative">
    <video width="320" height="240" ref="remoteVideo" autoplay></video>
    <video width="320" height="240" ref="localVideo" autoplay></video>
    <div class="buttons">
      <BaseButton @click="dialogState = true">Invite friend</BaseButton>
      <BaseButton @click="handleLeave" color="red">Leave</BaseButton>
    </div>
  </div>
  <BaseModal @close="dialogState = false" :dialog-state="dialogState">
    <template #title>Invite friend</template>
    <template #body>
      <BaseInput
        v-model="username"
        placeholder="Your name"
        style="width: 100%"
      />
      <BaseInput
        v-model="friendMail"
        placeholder="Your frinend's mail"
        style="width: 100%; margin-top: 10px"
      />
    </template>
    <template #actions>
      <BaseButton @click="invite">Invite</BaseButton>
    </template>
  </BaseModal>
</template>

<style scoped>
.container {
  display: flex;
  gap: 10px;
}
video {
  border: 1px solid #ccc;
}
.buttons {
  position: absolute;
  display: flex;
  gap: 10px;
  right: 0%;
  bottom: -17%;
}
</style>
