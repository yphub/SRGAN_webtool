<template>
  <div id="app">
    <Navbar />
    <Tool />
  </div>
</template>

<script>
import Navbar from "./components/Navbar.vue";
import Tool from "./components/Tool.vue";
export default {
  name: "app",
  components: {
    Navbar,
    Tool
  },
  data() {
    return {};
  },
  mounted() {},
  methods: {}
};
</script>

<style lang="scss">
body {
  margin: 0;
  user-select: none;
}
#app {
  position: absolute;
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
  .fade-enter-active,
  .fade-leave-active {
    transition: opacity ease-in-out 0.4s;
  }
  .fade-enter,
  .fade-leave-to {
    opacity: 0;
  }
}
</style>
