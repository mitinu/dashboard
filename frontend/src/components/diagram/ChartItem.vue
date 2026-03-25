<template>
  <div class="chartItem"

  >
    <svg
        ref="svg"
        :viewBox="`0 0 ${internalSizeWidthSvg} ${internalSizeHeightSvg}`"
        preserveAspectRatio="none"
        :style="{
          minWidth: `${widthColum}px`,
        }"

    >
      <g :transform="`translate(0, ${internalSizeHeightSvg}) scale(1, -1)`">
        <polygon
          v-if="lowerBackgroundFill"
          :points="fillPoints"
          :fill="lowerBackgroundFillColor"
        />
        <!-- Линия к предыдущей точке -->
        <line
          v-if="!isFirstPoint"
          :x1="internalSizeWidthSvg * 1.50"
          :y1="(previousValue - minValue)*k"
          :x2="internalSizeWidthSvg / 2"
          :y2="(value - minValue)*k"
          :stroke="color"
          stroke-width="1"
          stroke-linecap="round"
          stroke-linejoin="round"
        />

        <!-- Линия к следующей точке -->
        <line
          v-if="!isLastPoint"
          :x1="-internalSizeWidthSvg / 2"
          :y1="(followingValue - minValue)*k"
          :x2="internalSizeWidthSvg / 2"
          :y2="(value - minValue)*k"
          :stroke="color"
          stroke-width="1"
          stroke-linecap="round"
          stroke-linejoin="round"
        />

        <!-- Точка -->
        <circle
          :cx="internalSizeWidthSvg / 2"
          :cy="(value - minValue)*k"
          :r="2"
          :fill="color"
          :stroke="frameColor"
          stroke-width="0.5"
        />
      </g>
    </svg>
  </div>
</template>

<script>
export default {
  name:"chartItem",
  props:{
    widthColum:{type: Number},
    value:{type: Number},
    previousValue:{type: Number},
    followingValue:{type: Number},
    maxValue:{type: Number},
    minValue:{type: Number},
    color:{type: String},
    frameColor:{type: String, default:"#fbfbfb"},
    lowerBackgroundFill:{type: Boolean, default: false},
    lowerBackgroundFillColor:{type: String, default:"rgba(65,36,103,0.35)"}
  },
  data() {
    return{
      isFirstPoint: this.previousValue ? false:true,
      isLastPoint: this.followingValue ? false:true,
      containerWidth: 80,
      containerHeight: 500,
    }
  },
  computed: {
    k(){
      // return this.containerHeight/((this.maxValue-this.minValue)*3)
      // return this.containerHeight / ((this.maxValue - this.minValue) * Math.sqrt(this.containerHeight) / 3)
      // return (3 * Math.sqrt(this.containerHeight)) / (this.maxValue - this.minValue);
      return Math.sqrt(this.containerHeight) / (0.30 * (this.maxValue - this.minValue));
    },
    internalSizeHeightSvg() {
      return (this.maxValue-this.minValue)*this.k;
    },
    internalSizeWidthSvg() {
      return  this.containerWidth*(this.internalSizeHeightSvg / this.containerHeight);
    },
    fillPoints() {
      const centerX = this.internalSizeWidthSvg / 2;
      let points = [];

      // Начальная точка (слева внизу)
      points.push(`0,-500`);

      // Если это не последняя точка, добавляем правую линию
      if (!this.isLastPoint)  {
        const rightX = -this.internalSizeWidthSvg / 2
        const leftY = (this.followingValue - this.minValue) * this.k;
        points.push(`${rightX},${leftY}`);
      }
      else {
        const rightX = -this.internalSizeWidthSvg / 2
        const leftY = (this.value - this.minValue) * this.k;
        points.push(`${rightX},${leftY}`);
      }


      // Текущая точка
      points.push(`${centerX},${(this.value - this.minValue) * this.k}`);

      // Если это не первая точка, добавляем левую линию
      if (!this.isFirstPoint) {
        const rightX = this.internalSizeWidthSvg * 1.50;
        const rightY = (this.previousValue - this.minValue) * this.k;
        points.push(`${rightX},${rightY}`);
      }
      else {
        const rightX = this.internalSizeWidthSvg * 1.50;
        const rightY = (this.value - this.minValue) * this.k;
        points.push(`${rightX},${rightY}`);
      }

      points.push(`${this.internalSizeWidthSvg}, -500`);


      // Замыкающая точка (справа внизу)
      points.push(`${this.internalSizeWidthSvg}, -500`);

      return points.join(' ');
    }
  },
  mounted() {
    this.updateContainerHeight();
    window.addEventListener('resize', this.updateContainerHeight);

    this.updateContainerWidth();
    window.addEventListener('resize', this.updateContainerWidth);
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.updateContainerHeight);
    window.removeEventListener('resize', this.updateContainerWidth);
  },
  methods: {
    updateContainerHeight() {
      if (this.$refs.svg) {
        this.containerHeight = this.$refs.svg.clientHeight || 500;
      }
    },
    updateContainerWidth() {
      if (this.$refs.svg) {
        this.containerWidth = this.$refs.svg.clientWidth || 500;
      }
    },
  }
}
</script>

<style scoped>
.chartItem{
  height: 100%;
  width: 100%;
  svg {
    display: block;
    height: 100%;
    width: 100%;
  }
}
</style>