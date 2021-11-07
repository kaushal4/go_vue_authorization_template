<template>
  <div v-if="name == ''">loading</div>
  <div v-else>
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
      <h1>Welcome! {{ name }}</h1>
      <button type="button" class="btn btn-danger" @click="handleLogOut">
        Log Out
      </button>
      <Register :user="name" v-if="isTeacher" />
      <ViewCourse :type="type" :user="name" />
    </div>
  </div>
</template>

<script lang="ts">
import axios from "axios";
import { defineComponent } from "@vue/runtime-core";
import router from "@/router";
import ViewCourse from "../components/ViewCourse.vue";
import RegisterCourseVue from "@/components/RegisterCourse.vue";

export default defineComponent({
  name: "studentHomePage",
  props: {
    type: {
      type: String,
      required: true,
    },
  },
  components: {
    ViewCourse,
    Register: RegisterCourseVue,
  },
  data() {
    return {
      name: "" as string,
      courseId: "" as string,
    };
  },
  async beforeMount(): Promise<void> {
    try {
      const result = await axios.get(`http://localhost:8000/${this.type}`, {
        withCredentials: true,
      });
      this.name = result.data.user;
    } catch (err) {
      router.push(`/login/teacher`);
    }
  },
  methods: {
    async handleLogOut(): Promise<void> {
      const result = await axios.get(
        `http://localhost:8000/${this.type}/logout`,
        {
          withCredentials: true,
        }
      );
      if (result.status === 200) router.push("/");
    },
  },
  computed: {
    isTeacher(): boolean {
      return this.type === "teacher" ? true : false;
    },
  },
});
</script>
