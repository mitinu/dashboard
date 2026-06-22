<template>
  <div>
    <canvas
        ref="canvas"
        :style="{
        minWidth: `${widthColum}px`
      }"
    >
    </canvas>
  </div>
</template>

<script>

export default {
  props:{
    widthColum:{type: Number},
    values:{type: Array},
    maxValue:{type: Number},
    minValue:{type: Number},
    ratioFillingWidth:{type: Array, default:[2, 2]},
    colors:{type: Array}
  },
  name: "comparisonHistogramItem",
  data() {
    return{
      internalSizeHeightCanvas: this.maxValue-this.minValue,
      internalSizeWidthCanvas: (this.ratioFillingWidth[0]*10+this.ratioFillingWidth[1]*20)*this.values.length+this.ratioFillingWidth[0]*10
    }
  },
  mounted() {

    this.$refs.canvas.height
    const ctx = this.$refs.canvas.getContext('2d')



    // Вычисляем коэффициенты масштабирования
    const scaleX = this.$refs.canvas.width / this.internalSizeWidthCanvas;
    const scaleY = this.$refs.canvas.height / this.internalSizeHeightCanvas;



    // Применяем масштаб
    // Переворачиваем ось Y и смещаем начало координат
    ctx.translate(0, this.$refs.canvas.height)
    ctx.scale(scaleX, -scaleY);


    // Рисуем с координатами от левого нижнего угла
    let xPosition = 0
    for (const idx in this.values) {
      const value = this.values[idx]
      ctx.fillStyle = this.colors[idx].color
      ctx.fillRect(xPosition+this.ratioFillingWidth[0]*10, 0, this.ratioFillingWidth[1]*20, value-this.minValue)
      xPosition += this.ratioFillingWidth[0]*10 + this.ratioFillingWidth[1]*20
    }

  },
}
</script>

<style scoped>
canvas{
  height: 100%;
  width: 100%;
}
</style>