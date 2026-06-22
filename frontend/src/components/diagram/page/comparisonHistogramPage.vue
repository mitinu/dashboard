<template>
  <div class="histogramAndChartPage">
    <headerDiagramPage
        :legend="legend"
        :title="title"
    />
    <axesFrame
        :leftYAxis="leftYAxis"
        :XAxis="XAxis"
        :widthColum="widthColum"
    >
      <diagramGrid
          :numberSteps="rangeBoundariesHistogram.numberSteps"
          :dates="dates"
          :widthColum="widthColum"
      />
      <comparison-histogram
          v-if="dataset.histogram"
          :minValue="rangeBoundariesHistogram.minBorder"
          :maxValue="rangeBoundariesHistogram.maxBorder"
          :dataset="dataset.histogram"
          :dates="dates"
          :colors="dataset.histogramColor"
          :widthColum="widthColum"
          :config="dataset.histogramConfig"
      />
    </axesFrame>
  </div>
</template>

<script>
import HeaderDiagramPage from "@/components/diagram/page/component/headerDiagramPage.vue";
import diagramGrid from "@/components/diagram/DiagramGrid.vue";
import comparisonHistogram from "@/components/diagram/ComparisonHistogram.vue";
import AxesFrame from "@/components/diagram/page/component/AxesFrame.vue";

export default {
  components:{
    AxesFrame,
    comparisonHistogram,
    diagramGrid,
    HeaderDiagramPage
  },
  props:{
    title:{type: String},
    dataset:{type: Object},
    dates:{type: Array},
    legend:{type: Array}
  },
  name: "comparisonHistogramPage",
  data(){
    return{
      widthColum: 80,
      XAxis:{
        items: this.dates
      }
    }
  },
  computed:{
    leftYAxis() {
      const arr = []
      for (let idx = 0; idx < this.rangeBoundariesHistogram.numberSteps + 1; idx++) {
        arr.push(this.rangeBoundariesHistogram.maxBorder - idx * this.rangeBoundariesHistogram.sizeStep)
      }
      return {
        title: this.dataset.histogramYAxisTitle,
        items: arr
      }
    },
    rangeBoundariesHistogram(){
      //TODO это надо исправлять
      let maxBorder = ((obj = this.dataset.histogram) => {for (const key in obj) return obj[key][0];})();
      let minBorder = ((obj = this.dataset.histogram) => {for (const key in obj) return obj[key][0];})();

      if (minBorder<1)
      {
        maxBorder = this.roundingUp(maxBorder, 1, 2, this.dataset.histogram, 1.1);
        minBorder = this.roundingUp(minBorder, -1, 1, this.dataset.histogram);
      }
      else {
        maxBorder = this.roundingUp(maxBorder, 1, 1, this.dataset.histogram, 1.1);
        console.log(maxBorder)
        minBorder = this.roundingUp(minBorder, -1, 1, this.dataset.histogram);
      }
      let displayedRange = this.searchRange(maxBorder, minBorder)

      let numberSteps

      let lastDigit = this.lastDigit(displayedRange)

      switch (lastDigit){
        case 2:
        case 4:
        case 8:
          numberSteps = 8
          break
        case 1:
        case 3:
        case 9:
          numberSteps = 9
          break
        case 5:
          numberSteps = 5
          break
        case 6:
          numberSteps = 6
          break
        case 7:
          numberSteps = 7
          break
        default:
          numberSteps = displayedRange
          break
      }
      let x = parseInt(displayedRange.toString().slice(0,2))
      if (x%numberSteps!=0){
        displayedRange=parseInt(x+(numberSteps-x%numberSteps)+displayedRange.toString().slice(2))
      }
      let sizeStep = displayedRange/numberSteps
      if(minBorder%sizeStep!=0){
        minBorder = Math.floor(minBorder/sizeStep)*sizeStep
      }
      maxBorder = displayedRange+minBorder
      console.log(maxBorder)

      return {
        maxBorder: maxBorder,
        minBorder: minBorder,
        displayedRange: displayedRange,
        numberSteps: numberSteps,
        sizeStep: sizeStep
      }


    },
  },
  methods:{
    validationDisplayedRange(value) {
      const simpleNumber = [11, 13, 17, 19, 22, 23, 26, 29, 31, 33, 34, 37, 38, 39, 41, 43, 44, 46, 47, 51, 52, 53, 55, 57, 58,
        59, 61, 62, 65, 66, 67, 68, 69, 71, 73, 74, 76, 77, 78, 79, 82, 83, 85, 86, 87, 88, 89, 91, 92, 93, 94, 95, 97, 99]

      if (simpleNumber.includes(value.toString().slice(0, 2))) return false
      return true
    },

    roundingUp(startValue, bigWay, charactersBeforeRounding, arrs, roundingFactor = 1.0){
      let thisValue = startValue
      for (const i in arrs) {
        const arr = arrs[i]
        for (const idx in arr){
          const value = arr[idx]
          if (thisValue*bigWay<value*bigWay){
            thisValue=value
          }
        }
      }
      if(bigWay==1&&thisValue>0||bigWay==-1&&thisValue<0){
        thisValue = thisValue * roundingFactor
      }
      else {
        thisValue = thisValue / roundingFactor

      }

      if (Math.abs(thisValue) >= 10){
        const factor = Math.pow(10, Math.floor(Math.log10(Math.abs(Math.round(thisValue)))) - (charactersBeforeRounding-1));
        if(bigWay==1){
          thisValue = Math.ceil(thisValue / factor) * factor;
        }
        if(bigWay==-1){
          thisValue =  Math.floor(thisValue / factor) * factor;
        }

        return parseInt(thisValue)
      }
      else{
        if(bigWay==1){
          thisValue = this.roundUp(thisValue, 1)
        }
        if(bigWay==-1){
          thisValue = this.roundDown(thisValue, 1)

        }
        return parseFloat(thisValue)
      }
      return -1
    },
    lastDigit(n) {
      if(n<1){
        n *= 10
      }
      let num = n;
      let d = 2;

      while (num >= 10) {
        if (num % d === 0) {
          num = num / d;
        } else {
          d++;
        }
      }

      return num;
    },


    roundUp(value, charactersBeforeRounding = 0) {
      const multiplier = Math.pow(10, charactersBeforeRounding);
      return (Math.ceil(value * multiplier) / multiplier).toFixed(charactersBeforeRounding);
    },
    roundDown(value, charactersBeforeRounding = 0) {
      const multiplier = Math.pow(10, charactersBeforeRounding);
      return (Math.floor(value * multiplier) / multiplier).toFixed(charactersBeforeRounding);
    },
    formatValue(value) {
      return Number(value.toFixed(2));
    },
    searchRange(max, min){
      let range = Number((max - min).toFixed(10));
      if (range>10){
        if(!this.validationDisplayedRange(range)){
          range = (parseInt(range.toString().slice(0, 2))+1)*parseInt("1"+range.toString().slice(2))
        }
      }
      return range
    }
  }
}
</script>

<style scoped>
.histogramAndChartPage{
  height: 100%;
  padding: 20px 30px;
  box-sizing: border-box;
}
</style>