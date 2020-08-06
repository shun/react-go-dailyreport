<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">
        <span><h1>Daily report</h1></span>
      </div>
    </v-app-bar>

    <v-main>
      <div id="nav" style="margin: 10px;">
        <router-link to="/">日報登録</router-link> |
        <router-link to="/search">日報検索</router-link>
      </div>

      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "App",

  components: {},

  data: () => ({
    //
  })
});
</script>
