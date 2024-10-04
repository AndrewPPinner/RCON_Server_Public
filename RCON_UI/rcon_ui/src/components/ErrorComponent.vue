  <template>
  <transition-group class="alerts-drawer-wrapper" name="alerts-drawer-wrapper" tag="div">
    <template v-for="error in errors" :key="error[0]">
      <v-card
        class="alerts-snack rounded-lg elevation-8"
        maxWidth="500px"
        :color="'red'"
        :minWidth="'80%'">
        <v-card-actions class="pl-4 pr-5">
          <span class="grey--text text--lighten-3 body-1" v-text="error[1].Error"/>
          <v-icon @click="popError(error[0])" class="pl-4" icon="fa fa-times" />
        </v-card-actions>
      </v-card>
    </template>
  </transition-group>
</template>

<script lang="ts">
import store from '@/store';
import { defineComponent } from 'vue'

export default defineComponent({
  data() {
    return {
    }
  },
  computed: {
    errors() {
      return store.state.errors
    },
    hasErrors() {
      return store.state.errors.has(0)
    },
    timeOutErrors() {
      return true;
    }
  },
  methods: {
    popError(id : number) {
      store.dispatch("removeError", id)
    }
  },
  mounted() {
    setInterval(() => {
      for (let [key, value] of store.state.errors) {
        if (value.Time <= new Date(new Date().getTime())) {
          this.popError(key);
        }
      }
    }, 1000);
  },
})
</script>

<style>
.alerts-drawer-wrapper {
  max-width: 100%;
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: end;
  position: fixed;
  pointer-events: none;
  z-index: 1000;
  > .alerts-snack {
    margin-bottom: 2rem;
    pointer-events: auto;
  }
  &-enter, &-leave-to, &-leave-active {
    opacity: 0;
    transform: translateY(-30px);
  }
  &-leave-active{
    position: absolute
  }
}
</style>
