<template>
    <v-container>
        <v-row>
            <v-col>
            <v-sheet style="width: 500px;" class="bg-grey-darken-4 pa-12" rounded>
              <v-card class="mx-auto px-6 py-8">
                <v-btn
                  @click="openGarage()"
                  append-icon="fas fa-warehouse"
                  color="success"
                  size="large"
                  variant="elevated"
                  block>
                  Open Garage
                </v-btn>
              </v-card>
            </v-sheet>
          </v-col>
        </v-row>
    </v-container>
</template>
  
  <script lang="ts">
  import { defineComponent } from 'vue'
  import SmartHomeService from '@/services/SmartHomeService';
  
  export default defineComponent({
    data() {
      return {
      }
    },
    methods: {
        async openGarage() {
        await SmartHomeService.OpenGarage();
      },
    }
  })
  </script>
  