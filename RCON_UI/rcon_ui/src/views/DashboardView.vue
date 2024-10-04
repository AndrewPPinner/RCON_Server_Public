<template>
  <v-card class="bg-grey-darken-4">
    <v-tabs v-model="tab" align-tabs="center">
      <v-tab text="Palworld RCON" value="tab-1"></v-tab>
      <v-tab text="Smart Home" value="tab-2"></v-tab>
      <v-tab text="Minecraft" value="tab-3"></v-tab>
    </v-tabs>
    
    <v-window v-model="tab">
      <v-window-item value="tab-1"><PalWorldDashboard></PalWorldDashboard></v-window-item>
      <v-window-item value="tab-2"><SmartHomeDashboard></SmartHomeDashboard></v-window-item>
      <v-window-item value="tab-3"></v-window-item>
    </v-window>
  </v-card>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import PalWorldDashboard from '@/components/PalworldComponent.vue'
import SmartHomeDashboard from '@/components/SmartHomeComponent.vue'

export default defineComponent({
  components: {
    PalWorldDashboard,
    SmartHomeDashboard
  },
  data() {
    return {
      tab: 'tab-1',
    }
  }
  
})
</script>
