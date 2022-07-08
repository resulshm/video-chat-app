<script setup lang="ts">
import { RouterView } from "vue-router";
import TheHeader from "@/components/app/TheHeader.vue";
</script>

<template>
  <div>
    <header class="header">
      <TheHeader />
    </header>
    <main>
      <RouterView />
    </main>
  </div>
</template>

<style>
html {
  box-sizing: border-box;
}
*,
*:before,
*:after {
  box-sizing: inherit;
}

body {
  font-family: Arial, Helvetica, sans-serif;
}
header {
  padding: 0 15px;
  position: fixed;
  top: 0;
  left: 0;
}
main {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
</style>
