<template>
  <div class="axesPage">
    <div class="leftYAxis YAxis"
      v-if="leftYAxis"
    >
      <div class="headerDiagram">
        <span class="finePrint">{{ leftYAxis.title }}</span>
      </div>
      <div class="items">
        <div class="item"
             v-for="(item, idx) in leftYAxis.items"
             :key="idx"
        >
          <span class="finePrint">{{item}}</span>
        </div>
      </div>
    </div>
    <div class="content"
         ref="sliderRef"
         @mousedown="handleMouseDown"
         @mousemove="handleMouseMove"
         @mouseup="handleMouseUp"
         @mouseleave="handleMouseLeave"
    >
      <div class="headerDiagram"></div>
      <div
          class="diagramBody"
      >
        <slot></slot>
      </div>
      <div class="XAxis">
        <div class="container">
          <div class="finePrint"
               v-for="(item, idx) in XAxis.items"
               :key="idx"
               :style="{
                minWidth: `${widthColum}px`,
              }"
          >
            <span>{{item}}</span>
          </div>
        </div>
      </div>

    </div>
    <div class="rightYAxis YAxis"
      v-if="rightYAxis"
    >
      <div class="headerDiagram">
        <span class="finePrint">{{ rightYAxis.title }}</span>
      </div>
      <div class="items">
        <div class="item"
             v-for="(item, idx) in rightYAxis.items"
             :key="idx"
        >
          <span class="finePrint">{{ item }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>


export default {
  name: "axesFrame",
  props:{
    leftYAxis:{type: Object},
    XAxis:{type: Object},
    rightYAxis:{type: Object},
    widthColum:{type: String},
  },
  data(){
    return{
      isDown: false,
      startX: 0,
      scrollLeft: 0,
      velocity: 0,
      rafId: null,
    }
  },
  methods:{

    handleMouseDown(e) {
      if (e.button !== 0) return
      e.preventDefault()
      const slider = this.$refs.sliderRef
      this.isDown = true
      this.startX = e.pageX - slider.offsetLeft
      this.scrollLeft = slider.scrollLeft

      if (this.rafId) {
        cancelAnimationFrame(this.rafId)
      }
      this.velocity = 0

      slider.style.cursor = 'grabbing'
      e.preventDefault()
    },
    handleMouseMove(e) {
      if (!this.isDown) return
      e.preventDefault()

      const slider = this.$refs.sliderRef
      const x = e.pageX - slider.offsetLeft
      const walk = (x - this.startX) * 0.8
      const newScrollLeft = this.scrollLeft - walk

      this.velocity = newScrollLeft - slider.scrollLeft
      slider.scrollLeft = newScrollLeft
    },
    handleMouseUp() {
      const slider = this.$refs.sliderRef
      this.isDown = false
      slider.style.cursor = 'grab'

      if (Math.abs(this.velocity) > 0.5) {
        this.startInertia()
      }
    },
    handleMouseLeave() {
      const slider = this.$refs.sliderRef
      this.isDown = false
      slider.style.cursor = 'grab'
    },
    startInertia() {
      const slider = this.$refs.sliderRef

      const step = () => {
        this.velocity *= 0.92
        slider.scrollLeft += this.velocity

        if (Math.abs(this.velocity) > 0.1) {
          this.rafId = requestAnimationFrame(step)
        }
      }

      this.rafId = requestAnimationFrame(step)
    },
  },
  mounted() {
    this.$nextTick(() => {
      const slider = this.$refs.sliderRef
      if (slider) {
        // Устанавливаем скролл в конец
        slider.scrollLeft = slider.scrollWidth
      }
    })
  },
  beforeUnmount() {
    if (this.rafId) {
      cancelAnimationFrame(this.rafId)
    }
  }
}
</script>

<style scoped>
.axesPage{
  height: calc(100% - 100px);
  display: grid;

  grid-template-columns: 50px 1fr 50px;
  .headerDiagram{
    height: 35px;
    display: flex;
    justify-content: center;
    align-items: start;
  }
  .YAxis{
    height: 100%;
    background-color: #171620;
    .items{
      height: calc(100% - 35px);
      display: flex;
      flex-direction: column;
      .item{
        position: relative;
        flex: 1;
        min-height: 0;
        &:last-child {
          flex: 0 0 25px;
        }
        span{
          position: absolute;
          z-index: 5;
          top: -7px;
        }
      }
    }
  }
  .leftYAxis .items .item span{
    right: 15px;
  }
  .rightYAxis .items .item span{
    left: 15px;
  }
  .content{
    height: 100%;
    overflow-x: auto;
    overflow-y: hidden;
    white-space: nowrap;

    cursor: grab;
    cursor: -moz-grab;
    cursor: -webkit-grab;

    scroll-behavior: smooth;

    user-select: none;

    scrollbar-width: none;
    -ms-overflow-style: none;
    ::-webkit-scrollbar {
      display: none;
    }
    :active {
      cursor: grabbing;
      cursor: -moz-grabbing;
      cursor: -webkit-grabbing;
    }
    .diagramBody{
      height: calc(100% - 20px - 35px);
      position: relative;
    }
    .XAxis{
      position: relative;
      .container{
        width: 100%;
        position: absolute;
        z-index: 1;
        min-width: min-content;
        display: flex;
        flex-direction: row-reverse;
        height: 20px;
      }
      .finePrint{
        width: 100%;
        display: flex;
        justify-content: center;
      }
    }
  }
}
</style>