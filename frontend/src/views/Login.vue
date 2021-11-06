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
    <h1>{{ capitalizedUser }} Login</h1>
    <form v-on:submit.prevent="handleSubmit">
      <div class="mb-3">
        <label for="exampleInputEmail1" class="form-label">Name</label>
        <input
          type="text"
          class="form-control"
          aria-describedby="emailHelp"
          v-model="userName"
        />
      </div>
      <div class="mb-3">
        <label for="exampleInputPassword1" class="form-label">Password</label>
        <input type="password" class="form-control" v-model="password" />
      </div>
      <button type="submit" class="btn btn-primary">Submit</button>
    </form>
  </div>
</template>

<script lang="ts">
import router from "@/router";
import { defineComponent } from "@vue/runtime-core";
import axios from "axios";
export default defineComponent({
  props: {
    id: String,
  },
  beforeMount(): void {
    if (this.id != "student" && this.id != "teacher") {
      router.push("/");
    }
  },
  data() {
    return {
      userName: "" as string,
      password: "" as string,
    };
  },
  computed: {
    capitalizedUser(): string {
      if (typeof this.id === "string")
        return this.id.charAt(0).toUpperCase() + this.id.slice(1);
      else return "";
    },
  },
  methods: {
    async handleSubmit(): Promise<void> {
      interface user {
        name: string;
        password: string;
      }
      const currentDetails: user = {
        name: this.userName,
        password: this.password,
      };
      const result = await axios.get(`http://localhost:8000/${this.id}/login`, {
        params: currentDetails,
        withCredentials: true,
      });
      if (result.status == 200) {
        router.push(`/${this.id}`);
      }
    },
  },
});
</script>
