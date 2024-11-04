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
                Recorded Time: {{ sensor.CreatedAt }} <br/>
                <br/>
              </div>
            </v-card>
          </v-sheet>
        </v-col>
      </v-row>
      <v-sheet class="bg-grey-darken-4 pa-12" rounded>
            <v-card class="mx-auto px-6 py-8">
              <Line :data="chartJSGraph" :options="chartJSOptions"/>
            </v-card>
          </v-sheet>
    </v-container>
</template>
  
  <script lang="ts">
  import { defineComponent } from 'vue'
  import SmartHomeService from '@/services/SmartHomeService';
  import { Line } from 'vue-chartjs';
  import { CategoryScale, Chart, Legend, LinearScale, LineElement, PointElement, Title, Tooltip } from 'chart.js';

  Chart.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
  )
  
  export default defineComponent({
    components: {
      Line
    },
    data() {
      return {
        Sensors: [],
        graphData: [] as any[],
        graphLabels: [] as any[],
        chartJSGraph: {
          datasets: [] as any[],
          labels: [] as any[]
        },
        chartJSOptions: {
          responsive: true,
          maintainAspectRatio: false
        }
      }
    },
    methods: {
        async openGarage() {
          await SmartHomeService.OpenGarage();
        }
    },
    async mounted() {
      this.Sensors = await SmartHomeService.GetSensorValues();
      var data = await SmartHomeService.GetSensorDataGraph("Cilantro", "Moisture_Sensor");
      data.sort((a,b) =>  {
        return Date.parse(a.Date) - Date.parse(b.Date);
      });
      this.graphData = data.map(x => x.SensorValue);
      this.graphLabels = data.map(x => x.Date);
      this.chartJSGraph = {
        labels: this.graphLabels,
        datasets: [
          {
            label: 'Moisture_Sensor',
            backgroundColor: '#0000FF',
            data: this.graphData
          }
        ]
      }
    }
  })
  </script>
  