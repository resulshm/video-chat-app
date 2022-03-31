<script setup lang="ts">
import { RouterView } from "vue-router";
import TheNavbar from "@/components/AppComponents/TheNavbar.vue";
</script>

<template>
  <div class="w-screen h-screen px-20 relative">
    <div class="absolute top-0 left-0">
      <TheNavbar />
    </div>
    <div class="w-full h-full flex justify-center items-center">
      <RouterView />
    </div>
  </div>
</template>

<style>
body {
  font-family: Arial, Helvetica, sans-serif;
}
</style>
