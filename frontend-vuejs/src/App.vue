<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">
        <span><h1>日報登録</h1></span>
      </div>

    </v-app-bar>

    <v-main>
      <RegistryComp />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import RegistryComp from "./components/RegistryComp.vue";

export default Vue.extend({
  name: "App",

  components: {
    RegistryComp
  },

  data: () => ({
    //
  })
});
</script>
