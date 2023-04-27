<template>
  <div class="cardItem">
    <h4 class="itemTitle">{{ itemKey }}</h4>
    <span>
      {{ itemData }}
    </span>
  </div>
</template>
  
<style>
.cardItem {
  @apply w-64 p-4 mr-4 my-4 box-border bg-slate-100 rounded-lg h-40 hover:bg-cyan-100
}

.itemTitle{
  @apply text-base w-full text-center font-bold capitalize
}
</style>

<script  lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'ItemProduct',
  props: {
    itemData: {
      required: true
    },
    itemKey: {
      required: true
    }
  }
});
</script>
  