<template>
  <v-sheet class="bg-grey-darken-4 pa-12" rounded>
    <v-card class="mx-auto px-6 py-8" max-width="344">
      <v-form
        v-model="form"
        @submit.prevent="onSubmit()">
        <v-text-field
          v-model="loginForm.Username"
          :readonly="loading"
          :rules="[required]"
          class="mb-2"
          label="Username"
          clearable>
        </v-text-field>

        <v-text-field v-model="loginForm.Password"
          :readonly="loading"
          :rules="[required]"
          label="Password"
          placeholder="Enter your password"
          clearable type="Password">
        </v-text-field>

        <br>

        <v-btn
          :disabled="!form"
          :loading="loading"
          color="success"
          size="large"
          type="submit"
          variant="elevated"
          block>
          Sign In
        </v-btn>
      </v-form>
    </v-card>
  </v-sheet>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { LoginRequest } from "@/models/LoginRequest";
import store from '@/store';

export default defineComponent({
  methods: {
    onSubmit () {
        if (!this.form) {
          return;
        } 
        this.loading = true;
        this.handleLogin();
    },
    required (v : any) {
        return !!v || 'Field is required'
    },
    handleLogin() {
      store.dispatch("login", this.loginForm).then(
        () => {
          this.$router.push("/");
          this.loading = false;
        },
        error => {
          this.loginForm.Password = "";
          this.loading = false;
        }
      )
    }
},
data() {
  return {
    loginForm: new LoginRequest(),
    form: false,
    loading: false,
  }
}
})
</script>
