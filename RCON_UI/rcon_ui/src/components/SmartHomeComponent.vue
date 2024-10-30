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
          
          <v-col>
            <v-sheet style="width: 500px;" class="bg-grey-darken-4 pa-12" rounded>
              <v-card class="mx-auto px-6 py-8">
                <label>Sensors:</label>
                <div v-for="sensor in Sensors" :key="sensor.ID">
                  <br/>
                  Sensor Value: {{ sensor.SensorValue }} <br/>
                  Sensor Location: {{ sensor.SensorLocation }} <br/>
                  Sensor Type: {{ sensor.SensorType }} <br/>
                  <br/>
                </div>

                <div>Raw Data: {{ Sensors }}</div>
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
        Sensors: []
      }
    },
    methods: {
        async openGarage() {
          await SmartHomeService.OpenGarage();
      },
    },
    async mounted() {
      this.Sensors = await SmartHomeService.GetSensorValues();
      console.log(this.Sensors);
    }
  })
  </script>
  