<template>
  <div class="mb-3">
    <label for="formFile" class="form-label">Default file input example</label>
    <input class="form-control" type="file" id="formFile" ref="file" />
    <button
      class="btn btn-primary mt-3"
      @click="handleUpload"
      :disabled="isDisabled"
    >
      {{ status }}
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
import axios from "axios";

export default defineComponent({
  props: {
    courseName: {
      type: String,
      required: true,
    },
    updateCourse: {
      type: Function,
      required: true,
    },
  },
  beforeMount() {
    axios.defaults.withCredentials = true;
  },
  data() {
    return {
      status: "Upload" as string,
    };
  },
  methods: {
    toggleLoading() {
      this.status = this.status == "Upload" ? "Loading..." : "Upload";
    },
    async handleUpload(): Promise<void> {
      debugger;
      if (
        this.$refs.file instanceof HTMLInputElement &&
        this.$refs.file.files != null
      ) {
        const file = this.$refs.file.files[0];
        const formdata = new FormData();
        formdata.append("file", file);
        formdata.append("name", this.courseName);
        this.toggleLoading();
        const result = await axios.post(
          "http://localhost:8000/teacher/course/addFile",
          formdata,
          {
            headers: {
              "Content-Type": "multipart/form-data",
            },
          }
        );
        if (result.status == 200) this.updateCourse();
        this.toggleLoading();
      }
    },
  },
  computed: {
    isDisabled(): boolean {
      if (this.status == "Uploading....") return true;
      return false;
    },
  },
});
</script>
