<template>
  <div class="myTable">
    <div
        class="table"
        :style="{
        gridTemplateColumns: gridColumns
      }"
    >
      <div
          class="cell1 stickyHeader"
          v-for="column in headerTable"
          :key="column.key"
      >
        <span>{{column.value}}</span>
      </div>
      <template
          v-for="(rowTable, key, index) in bodyTable"
          :key="key"
      >
        <div
            class="cell"
            v-for="column in headerTable"
            :key="column.key"
            :class="['table-row', { 'table-row-last': index === Object.keys(bodyTable).length - 1}]"
        >
        <span>
          {{rowTable[column.key]}}
        </span>
        </div>
      </template>
    </div>
  </div>
</template>

<script>
export default {
  name: "myTable",
  props:{
    headerTable:{type:Array, required:true},
    bodyTable:{type:Object, required:true}
  },
  computed: {
    gridColumns() {
      return this.headerTable.map(col => col.width || '1fr').join(' ')
    }
  }
}
</script>

<style scoped lang="scss">
.myTable{
  position: relative;
  .table{
    width: 100%;
    position: absolute;
    max-height: 100%;
    top: 0;
    left: 0;
    overflow-y: scroll;
    display: grid;

    div{
      padding: 5px;
    }

    .stickyHeader{
      position: sticky;
      top: 0;
      font-weight: bold;
      z-index: 1;
      padding-bottom: 15px;
    }

    .table-row {
      border-bottom: 1px solid #e0e0e0;
      padding-top: 12px;
      padding-bottom: 12px;
      &-last {
        border-bottom: none;
      }
    }

  }
}
</style>