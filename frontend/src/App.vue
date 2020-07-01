<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
  export default {
    watch: {
      $route: (to, from) => {
        document.title = to.meta.title || 'OPNV-Smartcity';
      },
    },
  };
</script>
