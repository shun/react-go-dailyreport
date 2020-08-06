import Vue from "vue";
import App from "./App.vue";
import router from "./router";
//import vuetify from "./plugins/vuetify";
import Vuetify from "vuetify";
//import '@fortawesome/fontawesome-free/css/all.css';
//import { library } from "@fortawesome/fontawesome-svg-core";
//import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
//import { fas } from "@fortawesome/free-solid-svg-icons";
import "material-design-icons-iconfont/dist/material-design-icons.css";

//Vue.component("font-awesome-icon", FontAwesomeIcon); // Register component globally
//library.add(fas); // Include needed icons

Vue.config.productionTip = false;
Vue.use(Vuetify);

new Vue({
  router,
  vuetify: new Vuetify({
    icons: {
      iconfont: "md"
    }
  }),
  render: h => h(App)
}).$mount("#app");
