<template>
  <nav class="navbar">
    <router-link to="/" class="navTitle">Home</router-link> 
  </nav>
  <router-view/>
</template>

<style>
#app {
  width: 100vw;
  height: 100vh;
}

.navbar {
  @apply  bg-sky-500/75 text-white py-3 pl-10
}

.navTitle {
  @apply text-3xl
}
</style>
