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
    value:{type: Number},
    maxValue:{type: Number},
    minValue:{type: Number},
    ratioFillingWidth:{type: Array, default:[1, 2]},
    color:{type: String}
  },
  name: "HistogramItem",
  data() {
    return{
      internalSizeHeightCanvas: this.maxValue-this.minValue,
      internalSizeWidthCanvas: this.ratioFillingWidth[0]*20+this.ratioFillingWidth[1]*20
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
    ctx.fillStyle = this.color

    ctx.fillRect(this.ratioFillingWidth[0]*10, 0, this.ratioFillingWidth[1]*20, this.value-this.minValue) // Квадрат на 50 пикселей выше нижнего края
  },
}
</script>

<style scoped>
canvas{
  height: 100%;
  width: 100%;
}
</style>