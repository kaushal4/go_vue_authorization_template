<template>
  <div class="container">
    <div class="row">
      <h3>Search for courses</h3>
    </div>
    <form class="row justify-content-center" @submit.prevent="handleSearch">
      <input type="text" v-model="courseId" class="col-2" />
      <div class="col-3">
        <button type="submit" class="btn btn-info mx-5">Get Course</button>
      </div>
    </form>
    <div v-if="isCourseSet">
      <h4>Course Name: {{ searchedCourse.name }}</h4>
      <ul class="list-group">
        <li
          v-for="(material, index) in searchedCourse.material"
          :key="index"
          class="list-group-item"
        >
          {{ material }}
          <button
            class="float-end btn btn-warning"
            v-if="shouldEdit"
            type="button"
            data-bs-toggle="offcanvas"
            data-bs-target="#offcanvasTop"
            aria-controls="offcanvasTop"
          >
            Edit
          </button>
          <div
            class="offcanvas offcanvas-top offcanvas-container"
            tabindex="-1"
            id="offcanvasTop"
            aria-labelledby="offcanvasTopLabel"
          >
            <div class="offcanvas-header">
              <h5 id="offcanvasTopLabel">Edit material</h5>
              <button
                type="button"
                class="btn-close text-reset mt-1"
                data-bs-dismiss="offcanvas"
                aria-label="Close"
                id="tempSolution"
              ></button>
            </div>
            <div class="offcanvas-body">
              <form @submit.prevent="handleEdit(index)">
                <div class="input-group">
                  <span class="input-group-text">With textarea</span>
                  <textarea
                    class="form-control"
                    aria-label="With textarea"
                    v-model="editMaterial"
                  ></textarea>
                </div>
                <button class="btn btn-warning mt-1">Edit</button>
              </form>
            </div>
          </div>
        </li>
      </ul>
      <ul class="list-group">
        <li
          v-for="(file, index) in pathName"
          :key="index"
          class="list-group-item"
        >
          {{ file }}
          <button
            class="float-end btn btn-warning"
            type="button"
            @click="handleDownload(index)"
          >
            <img
              src="https://img.icons8.com/material-outlined/24/000000/download--v1.png"
            />
          </button>
        </li>
      </ul>
      <Addfile
        v-if="shouldEdit"
        :courseName="searchedCourse.name"
        :updateCourse="handleSearch"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
import axios from "axios";
import Addfile from "./AddFile.vue";
interface course {
  name: string;
  material: string[];
  teacher: string;
  files: string[] | null;
}

export default defineComponent({
  props: {
    type: {
      type: String,
      required: true,
    },
    user: {
      type: String,
      required: true,
    },
  },
  components: {
    Addfile,
  },
  data() {
    return {
      searchedCourse: {} as course,
      courseId: "" as string,
      editMaterial: "" as string,
    };
  },
  beforeMount() {
    axios.defaults.withCredentials = true;
  },
  computed: {
    isCourseSet(): boolean {
      if (this.searchedCourse.name != null) return true;
      else return false;
    },
    shouldEdit(): boolean {
      if (this.searchedCourse.teacher == this.user && this.type === "teacher")
        return true;
      else return false;
    },
    pathName(): string[] | null {
      return (
        this.searchedCourse.files &&
        this.searchedCourse.files.map((element): string => {
          debugger;
          return element.substr(element.lastIndexOf("\\") + 1);
        })
      );
    },
  },
  methods: {
    async handleEdit(materialNo: number): Promise<void> {
      interface editBody {
        name: string;
        materialNo: string;
        material: string;
      }
      const body: editBody = {
        name: this.searchedCourse.name,
        materialNo: `${materialNo}`,
        material: this.editMaterial,
      };
      const result = await axios.patch(
        "http://localhost:8000/teacher/course",
        body
      );
      if (result.status === 200) {
        setTimeout(this.clickClose, 1000);
        this.searchedCourse.material[materialNo] = this.editMaterial;
      }
    },
    clickClose() {
      (document.querySelector("#tempSolution") as HTMLElement).click();
    },
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
      debugger;
      this.searchedCourse = result.data;
    },
    async handleDownload(index: number): Promise<void> {
      if (this.searchedCourse.files) {
        const response = await axios.get(`http://localhost:8000/file`, {
          withCredentials: true,
          responseType: "blob",
          params: {
            path: this.searchedCourse.files[index],
          },
        });
        const url = window.URL.createObjectURL(new Blob([response.data]));
        const link = document.createElement("a");
        link.href = url;
        link.setAttribute("download", "file.pdf"); //or any other extension
        document.body.appendChild(link);
        link.click();
      }
    },
  },
});
</script>

<style scoped>
.offcanvas-container {
  height: 40vh;
}
</style>
