<script lang="ts">
export default {
  name: "HomeView",
};
</script>

<script setup lang="ts">
import { ref } from "vue";
import { createRoom } from "@/api/room";
import IconChevronLeft from "../components/Icons/IconChevronLeft.vue";
import IconChevronRight from "../components/Icons/IconChevronRight.vue";
import IconVideo from "@/components/Icons/IconVideo.vue";
import BaseInput from "@/components/BaseComponents/BaseInput.vue";
import BaseButton from "@/components/BaseComponents/BaseButton.vue";
import { useRouter } from "vue-router";
const room = ref<string>();
const router = useRouter();
async function onSubmit() {
  const formData = new FormData();
  if (room.value) {
    formData.append("name", room.value);
  }
  const response = await createRoom(formData);
  if (response.success) {
    router.push(`/${response.data.room_id}`);
  }
}
</script>

<template>
  <div class="w-full h-screen flex flex-row justify-center items-center">
    <div class="w-1/2 px-20">
      <div class="text-[2.9rem] leading-[1.2] mb-6">
        Premium video meetings <br />
        Now free for everyone
      </div>
      <div class="text-[1.1rem] text-gray-600 mb-13">
        We re-enginereed the service we built for secure business meetings,<br />
        Google meet, to make it free and available for all
      </div>
      <form class="mb-8" @submit.prevent="onSubmit">
        <div class="w-full h-full flex flex-row">
          <div class="mr-7">
            <BaseButton>
              <IconVideo class="mr-3" />
              New meeting</BaseButton
            >
          </div>
          <div class="w-65">
            <BaseInput v-model="room" placeholder="Enter code or link" />
          </div>
        </div>
      </form>
      <div class="w-150 mb-8">
        <hr />
      </div>
      <div>
        <a class="text-c-blue-100" href="google.com">Learn more</a>
        about Google meet
      </div>
    </div>
    <div class="w-1/2 lg:h-screen flex justify-center items-center">
      <div class="relative">
        <div class="absolute left-[-10%] top-[50%]">
          <IconChevronLeft />
        </div>
        <div class="absolute right-[-10%] top-[50%]">
          <IconChevronRight />
        </div>
        <img
          class="rounded-full w-110 h-90"
          src="@/assets/img/computer-g51ba141c6_1920.jpg"
        />
        <div
          class="absolute w-full bottom-[-30%] flex justify-center items-center flex-col"
        >
          <p class="text-[1.5rem] font-normal">Get a link you can share</p>
          <p>Click <span>New meeting</span> to get a link you can send to</p>
          <p>people you want meet them</p>
        </div>
      </div>
    </div>
  </div>
</template>
