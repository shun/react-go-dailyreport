<template>
  <div class="search">
    <SearchComp />
  </div>
</template>

<script>
// @ is an alias to /src
import SearchComp from "@/components/SearchComp.vue";

export default {
  name: "Search",
  components: {
    SearchComp
  }
};
</script>
