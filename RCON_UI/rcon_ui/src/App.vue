<template>
  <div class="bg-grey-darken-4">
    <ErrorComponent v-if="hasErrors" />
    <router-view />
  </div>
</template>

<script lang="ts">
import ErrorComponent from '@/components/ErrorComponent.vue';
import { defineComponent } from 'vue';
import store from './store';

export default defineComponent({
  components: {
    ErrorComponent
  },
  data() {
    return {
      hasErrors: false,
    }
  },
  mounted() {
    store.watch(
    (state, getters) => state.errors && state.errors.size > 0,
    (val, oldVal) => {
      if (val) {
        if (store.state.errors != null){
          this.hasErrors = true;
        }
      } else {
        this.hasErrors = false;
      }
    }
  )
  },
})
</script>


<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

body {
  background-color: #212121;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>

<style>
  @import "https://use.fontawesome.com/releases/v5.0.13/css/all.css";
</style>
