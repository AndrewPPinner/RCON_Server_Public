<template>
    <div>   
      <v-container>
        <v-row no-gutters>
          <v-col>
            <v-sheet style="width: 500px;" class="bg-grey-darken-4 pa-12" rounded>
              <v-card class="mx-auto px-6 py-8" max-width="344">
                <v-text-field v-model="message"
                  label="Message to Broadcast"
                  placeholder="Broadcast Message"
                  clearable type="text">
                </v-text-field>
  
                <br>
  
                <v-btn
                  @click="broadcastMessage()"
                  :disabled="!message"
                  color="success"
                  size="large"
                  variant="elevated"
                  block>
                  Send Message
                </v-btn>
              </v-card>
            </v-sheet>
          </v-col>
  
          <v-col>
            <v-sheet style="width: 700px;" class="bg-grey-darken-4 pa-12" rounded>
              <v-card class="mx-auto px-6 py-8">
                <v-table style="min-height: 150px;">
                  <thead>
                    <tr>
                      <th class="text-left">Name</th>
                      <th class="text-left">UID</th>
                      <th class="text-left">SteamID</th>
                      <th class="text-left">
                        <v-btn
                          @click="refreshPlayers()"
                          append-icon="fas fa-sync-alt"
                          color="success"
                          size="small"
                          variant="elevated"
                          block>
                          Refresh
                        </v-btn>
                      </th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-show="loading">
                      <v-progress-circular
                        style="margin: 2rem;"
                        color="primary"
                        indeterminate>
                      </v-progress-circular>
                      <td>Loading Players...</td>
                    </tr>
                    <tr v-show="players?.length == 0 && !loading">
                      <td>No players online.</td>
                    </tr>
                    <tr v-for="player in players" :key="player.Name">
                      <td>{{ player.Name }}</td>
                      <td>{{ player.UID }}</td>
                      <td>{{ player.SteamID }}</td>
                      <td><button style="color: red;">Kick Me</button></td>
                    </tr>
                  </tbody>
                </v-table>
              </v-card>
            </v-sheet>
          </v-col>
  
          <v-col>
            <v-sheet style="width: 500px;" class="bg-grey-darken-4 pa-12" rounded>
              <v-card class="mx-auto px-6 py-8">
                <v-btn
                  @click="saveAndRestart()"
                  append-icon="fas fa-save"
                  color="success"
                  size="large"
                  variant="elevated"
                  block>
                  Save and Restart Server
                </v-btn>
              </v-card>
            </v-sheet>
          </v-col>
  
          <v-col>
            <v-sheet style="width: 500px;" class="bg-grey-darken-4 pa-12" rounded>
              <v-card class="mx-auto px-6 py-8">
                <label>Server Status:</label>
                <v-sheet :color="serverStatus == 'UP' ? 'success' : 'error'" :elevation="3" style="margin-bottom: 2rem;" class="mx-auto px-6 py-3">{{ serverStatus }}</v-sheet>
                <v-btn
                  @click="getStatus()"
                  append-icon="fas fa-sync-alt"
                  color="info"
                  size="large"
                  variant="elevated"
                  block>
                  Refresh
                </v-btn>
              </v-card>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>  
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue'
  import AdminCalloutService from '@/services/AdminCalloutService';
  import { Players } from '@/models/Players';
  
  export default defineComponent({
    data() {
      return {
        message: "",
        players: [] as Players[],
        loading: false,
        serverStatus: null
      }
    },
    async mounted() {
      this.loading = true;
      this.players = await AdminCalloutService.GetPlayers()
      this.loading = false;
      this.serverStatus = await AdminCalloutService.GetStatus();
    },
    methods: {
      broadcastMessage() {
        AdminCalloutService.BroadCastMessage(this.message);
        this.message = "";
      },
      async refreshPlayers() {
        this.loading = true;
        this.players = await AdminCalloutService.GetPlayers();
        this.loading = false;
      },
      async saveAndRestart() {
        await AdminCalloutService.SaveAndRestart();
      },
      async getStatus() {
        this.serverStatus = await AdminCalloutService.GetStatus();
      }
    }
  })
  </script>
  