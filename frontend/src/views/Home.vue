<template>
  <HelloWorld />
</template>

<script lang="ts">
import HelloWorldVue from "@/components/HelloWorld.vue";
import { defineComponent } from "@vue/runtime-core";
export default defineComponent({
  name: "home",
  components: {
    HelloWorld: HelloWorldVue,
  },
});
</script>

<style>
body {
  min-height: 100vh;
  background-color: #7f5a83;
  background-image: linear-gradient(315deg, #7f5a83 0%, #0d324d 74%);
  background-repeat: no-repeat;
}
</style>
