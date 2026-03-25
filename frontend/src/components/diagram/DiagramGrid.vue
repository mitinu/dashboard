<template>
  <div
      class="diagramGrid"
      :style="{
        gridTemplateColumns: generateGridTemplateColumns(),
        gridTemplateRows: `repeat(${numberSteps}, 1fr)`
      }"
  >
    <div
        class="cell"
        v-for='idx in numberSteps * (dates.length + (gridPositions == "onGridLine" ? 1 : 0))'
        :key="idx"
        :class="{ 'cell-small': isSmallCell(idx) }"
        :style="{
          borderLeft: verticalGrid,
        }"
    >
    </div>
  </div>
</template>

<script>
export default {
  name: "diagramGrid",
  props: {
    widthColum: { type: Number },
    numberSteps: { type: Number },
    dates: { type: Array },
    gridPositions:{type:String, default:"center"},
    verticalGrid:{type: String, default: "2px #363a45 dashed"}
  },
  methods: {
    generateGridTemplateColumns() {
      // Создаем колонки: dates.length + 1 (добавляем одну колонку)
      const totalColumns = this.dates.length;
      const columns = [];
      switch (this.gridPositions){
        case "onGridLine":
          for (let i = 0; i < totalColumns + 1; i++) {
            // Если это первая или последняя колонка
            if (i === 0 || i === totalColumns) {
              columns.push(`minmax(${this.widthColum / 2}px, 0.5fr)`);
            } else {
              columns.push(`minmax(${this.widthColum}px, 1fr)`);
            }
          }
          break
        case "center":
        default:
          for (let i = 0; i < totalColumns; i++) {
            columns.push(`minmax(${this.widthColum}px, 1fr)`);
          }
          break
      }
      return columns.join(' ');
    },
    isSmallCell(idx) {
      // Определяем индекс колонки для ячейки (0-based)
      const totalColumns = this.dates.length + 1;
      const columnIndex = (idx - 1) % totalColumns;
      // Возвращаем true если это первая или последняя колонка
      return columnIndex === 0 || columnIndex === totalColumns - 1;
    }
  }
}
</script>

<style scoped lang="scss">
.diagramGrid {
  width: 100%;
  position: absolute;
  z-index: 1;
  height: 100%;
  min-width: min-content;
  display: grid;

  .cell {
    border-top: 2px #363a45 dashed;
    &-small{

    }
  }


}
</style>