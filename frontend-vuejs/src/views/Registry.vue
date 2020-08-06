<template>
  <div class="registry">
    <RegistryComp />
  </div>
</template>

<script>
// @ is an alias to /src
import RegistryComp from "@/components/RegistryComp.vue";

export default {
  name: "Registry",
  components: {
    RegistryComp
  }
};
</script>
