<template>
  <div class="container">
    <div class="row">
      <div class="col-6 text-center">
        <h5>Clink to register course</h5>
      </div>
      <div class="col-6">
        <button class="btn btn-info" @click="setRegister">
          Register Course
        </button>
      </div>
    </div>
    <form v-if="isRegister" @submit.prevent="validateForm">
      <div class="input-group mb-3">
        <span class="input-group-text" id="basic-addon1">Name</span>
        <input
          type="text"
          class="form-control"
          placeholder="Course Name"
          aria-label="Username"
          aria-describedby="basic-addon1"
          v-model="courseId"
        />
      </div>
      <div>
        <div class="input-group input-group-sm mb-3">
          <span class="input-group-text" id="inputGroup-sizing-sm"
            >Number of Materials</span
          >
          <input
            type="number"
            class="form-control"
            aria-label="Sizing example input"
            aria-describedby="inputGroup-sizing-sm"
            v-model="materialNum"
            min="1"
            max="5"
          />
        </div>
      </div>
      <div class="input-group" v-for="index in materialNum" :key="index">
        <span class="input-group-text">With textarea</span>
        <textarea
          class="form-control"
          aria-label="With textarea"
          v-model="materials[index - 1]"
        ></textarea>
      </div>
      <div class="text-center">
        <button type="submit" class="btn btn-success">Sumbmit</button>
      </div>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
import axios from "axios";
export default defineComponent({
  props: {
    user: {
      type: String,
      require: true,
    },
  },
  beforeMount() {
    axios.defaults.withCredentials = true;
  },
  data() {
    return {
      isRegister: false as boolean,
      courseId: "" as string,
      materialNum: 1 as number,
      materials: [] as string[],
    };
  },
  methods: {
    validateForm(): void {
      if (this.courseId.length > 0 && this.materials[0].length > 0)
        this.handleSubmit();
    },
    async handleSubmit(): Promise<void> {
      const result = await axios.post("http://localhost:8000/teacher/course", {
        name: this.courseId,
        material: this.materials,
        teacher: this.user,
      });
      debugger;
    },
    setRegister(): void {
      this.isRegister = !this.isRegister;
    },
  },
});
</script>
