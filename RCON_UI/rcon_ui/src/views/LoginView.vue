<template>
  <div class="home">
    <LoginComponent />
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import LoginComponent from "@/components/LoginComponent.vue";

@Options({
  components: {
    LoginComponent,
  },
})
export default class LoginView extends Vue {}
</script>
