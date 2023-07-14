<template>
  <div class="searching-container">
    <form>
      <input type="text" v-model="searchValue" placeholder="Search" class="searching-input">
      <button class="searching-button">Search</button>
    </form>
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

.searching-container {
  @apply p-4
}

.searching-button {
  @apply border-solid border-2 border-sky-500 rounded-lg p-2 bg-sky-500 text-white hover:bg-white hover:text-black
}

.searching-input {
  @apply border-solid border-2 border-sky-500 p-2 border-solid border-2 border-sky-500 rounded-lg mx-4 w-80
}
</style>
  
<script lang="ts">
import { defineComponent, h, ref } from 'vue';
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
    const items = ref<unknown[]>([]);

    axios.get('http://localhost:3000')
      .then(response => {
        const data = response.data.hits.hits;
        console.log(response.data);

        let arrayData = data.map((element: { _source: any; }) => element._source);
         items.value = arrayData.map((element: any) => element);
         console.log(arrayData)
      })
      .catch(error => {
        console.error(error);
      });

    return {
      items,
    };
  },
  searchingForId(item: { nombre: string; }) {
    return item.nombre.toLowerCase().includes(this.searchValue.toLowerCase());
  }
});
</script>