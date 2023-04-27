<template>
  <ul class="responsive">
    <li v-for="(itemData, key) in items" :key="key">
      <item :itemKey="key" :itemData="itemData" />
    </li>
  </ul>
</template>

<style>
.responsive {
  @apply flex flex-wrap justify-center w-full py-6
}
</style>
  
<script lang="ts">
import { defineComponent, ref } from 'vue';
import Item from './Item.vue';
import axios from 'axios';

export default defineComponent({
  name: 'ListItems',
  components: {
    Item,
  },
  setup() {
    const items = ref([]);

    axios.get('http://localhost:3000/database')
      .then(response => {
        items.value = response.data.enron_mail.mappings.properties;
      })
      .catch(error => {
        console.error(error);
      });

    return {
      items,
    };
  },
});
</script>