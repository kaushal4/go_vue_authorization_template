<template>
  <div
    class="
      container
      mx-auto
      mt-5
      text-center
      bg-light bg-gradient
      p-5
      border border-dark border-3
      rounded
    "
  >
    <h1 class="p-5 pt-3">ISA ASSIGNMENT</h1>
    <div class="container">
      <div class="row">
        <div class="col m-3 h6">Name : Kaushal Damania</div>
        <div class="col m-3 h6">RegNo : 19BCE0156</div>
      </div>
      <div class="row">
        <div class="col m-3 h6">Professor : Amutha Prabhakar M.</div>
        <div class="col m-3 h6">Course Code : CSE:3501</div>
      </div>
      <div class="row">
        <div class="linkContainers">
          <router-link to="/login/student">Student Login</router-link>
          <router-link to="/login/teacher">Teacher Login</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";

export default defineComponent({
  name: "HelloWorld",
});
</script>

<style scoped>
.linkContainers {
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-template-rows: repeat(1fr);
}
</style>
