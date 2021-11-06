<template>
  <div class="container">
    <div class="row">
      <h3>Search for courses</h3>
    </div>
    <form class="row justify-content-center">
      <input type="text" v-model="courseId" class="col-2" />
      <div class="col-3">
        <button type="button" class="btn btn-info mx-5" @click="handleSearch">
          Get Course
        </button>
      </div>
    </form>
    <div v-if="isCourseSet">
      <h4>Course Name: {{ searchedCourse.name }}</h4>
      <ul class="list-group">
        <li
          v-for="material in searchedCourse.material"
          :key="material"
          class="list-group-item"
        >
          {{ material }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
import axios from "axios";
interface course {
  name: string;
  material: string[];
  teacher: string;
}

export default defineComponent({
  props: {
    type: {
      type: String,
      required: true,
    },
  },

  data() {
    return {
      searchedCourse: {} as course,
      courseId: "" as string,
    };
  },
  computed: {
    isCourseSet(): boolean {
      if (this.searchedCourse.name != null) return true;
      else return false;
    },
  },
  methods: {
    async handleSearch(): Promise<void> {
      const result = await axios.get(
        `http://localhost:8000/${this.type}/course`,
        {
          withCredentials: true,
          params: {
            name: this.courseId,
          },
        }
      );

      this.searchedCourse = result.data;
    },
  },
});
</script>
