<template>
  <div>
    <input type="text" v-model="searchValue" placeholder="Search">
    <button>Search</button>
  </div>
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
  data() {
    return {
      searchValue: '',
    };
  },
  setup() {
    const items = ref([]);

    axios.get('http://localhost:3000')
      .then(response => {
        // items.value = response.datas.hits.hits[0]._source;
        const hits = response.data.hits.hits;
        for (const hit of hits) {
          items.value = hit._source;
        }
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