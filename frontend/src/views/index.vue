<script lang="ts">
export default {
  name: "HomeView",
};
</script>

<script setup lang="ts">
import { ref } from "vue";
import { createRoom } from "@/api/room";
import IconChevronLeft from "@/components/icons/IconChevronLeft.vue";
import IconChevronRight from "@/components/icons/IconChevronRight.vue";
import IconVideo from "@/components/icons/IconVideo.vue";
import BaseInput from "@/components/base/BaseInput.vue";
import BaseButton from "@/components/base/BaseButton.vue";
import { useRouter } from "vue-router";
const meetingCode = ref<string>();
const router = useRouter();
async function onSubmit() {
  const formData = new FormData();
  formData.append("name", "New meet");
  try {
    const response = await createRoom(formData);
    if (response.success) {
      router.push(`/${response.data.room_id}`);
    }
  } catch (err) {
    console.error(err);
  }
}
function join() {
  router.push(`/${meetingCode.value}`);
}
</script>

<template>
  <div
    class="w-full h-full flex 2xl:flex-row flex-col justify-center items-center"
  >
    <div class="2xl:w-1/2 w-full 2xl:px-20 pt-23 max-2xl:text-center">
      <div class="text-[2.9rem] leading-[1.2] mb-6">
        Premium video meetings <br />
        Now free for everyone
      </div>
      <div class="text-[1.1rem] text-gray-600 mb-13">
        We re-enginereed the service we built for secure business meetings,<br />
        R&S meet, to make it free and available for all
      </div>
      <form class="mb-8" @submit.prevent="onSubmit">
        <div
          class="w-full h-full flex max-2xl:justify-center max-2xl:items-center"
        >
          <div class="mr-7">
            <BaseButton>
              <IconVideo class="mr-3" />
              New meeting</BaseButton
            >
          </div>
          <div class="w-65">
            <BaseInput v-model="meetingCode" placeholder="Enter code" />
          </div>
          <div
            @click="join"
            class="my-auto mx-3 text-c-blue-100 cursor-pointer"
          >
            Join
          </div>
        </div>
      </form>
      <div class="max-2xl:hidden w-150 mb-8">
        <hr />
      </div>
      <div class="max-2xl:hidden">
        <a
          class="text-c-blue-100"
          href="https://github.com/ResulShamuhammedov/video-chat-app"
          >Learn more</a
        >
        about R&S meet
      </div>
    </div>
    <div class="2xl:w-1/2 w-full 2xl:h-screen flex justify-center items-center">
      <div class="relative">
        <div
          class="absolute left-[-10%] top-[50%] active:bg-gray-300 rounded-full p-1 cursor-pointer transition-all duration-100 ease-in-out"
        >
          <IconChevronLeft />
        </div>
        <div
          class="absolute right-[-10%] top-[50%] active:bg-gray-300 rounded-full p-1 cursor-pointer transition-all duration-100 ease-in-out"
        >
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
          <p>
            Click <span class="font-bold">New meeting</span> to get a link you
            can send to
          </p>
          <p>people you want meet them</p>
        </div>
      </div>
    </div>
  </div>
</template>
