<template>
  <div class="fourChartPage">
    <headerDiagramPage
        :legend="legend"
        :title="title"
    />
    <axesFrame
        :leftYAxis="leftYAxis"
        :rightYAxis="dataset.charts2 ? rightYAxis : undefined"
        :XAxis="XAxis"
        :widthColum="widthColum"
    >
      <diagramGrid
          :numberSteps="rangeBoundariesChart.numberSteps"
          :dates="dates"
          :widthColum="widthColum"
          gridPositions="onGridLine"
          verticalGrid="none"
      />
      <template v-for="(chart, idx) in dataset.charts">
        <chart
            v-if="chart"
            :minValue="rangeBoundariesChart.minBorder"
            :maxValue="rangeBoundariesChart.maxBorder"
            :dataset="chart"
            :dates="dates"
            :color="dataset.chartsColor[idx].color"
            :widthColum="widthColum"
            :config="dataset.chartConfig"
        />
      </template>
      <template v-for="(chart, idx) in dataset.charts2">
        <chart
            v-if="chart"
            :minValue="rangeBoundariesChart2.minBorder"
            :maxValue="rangeBoundariesChart2.maxBorder"
            :dataset="chart"
            :dates="dates"
            :color="dataset.chartsColor2[idx].color"
            :widthColum="widthColum"
            :config="dataset.chartConfig"
        />
      </template>
    </axesFrame>
  </div>
</template>

<script>
import HeaderDiagramPage from "@/components/diagram/page/component/headerDiagramPage.vue";
import AxesFrame from "@/components/diagram/page/component/AxesFrame.vue";
import chart from "@/components/diagram/Chart.vue";
import diagramGrid from "@/components/diagram/DiagramGrid.vue";

export default {
  name: "fourChartPage",
  components: {
    diagramGrid,
    chart,
    AxesFrame,
    HeaderDiagramPage,
  },
  props:{
    title:{type: String},
    dataset:{type: Object},
    dates:{type: Array},
    legend:{type: Array}
  },
  data(){
    return{
      widthColum: 100,
      XAxis:{
        items: this.dates
      }
    }
  },
  computed:{
    leftYAxis() {
      const arr = []
      for (let idx = 0; idx < this.rangeBoundariesChart.numberSteps+1; idx++) {
        arr.push(this.rangeBoundariesChart.maxBorder - (idx) * this.rangeBoundariesChart.sizeStep)
      }
      return {
        title: this.dataset.chartYAxisTitle,
        items: arr
      }
    },
    rightYAxis() {
      const arr = []
      for (let idx = 0; idx < this.rangeBoundariesChart2.numberSteps+1; idx++) {
        arr.push(this.rangeBoundariesChart2.maxBorder - (idx) * this.rangeBoundariesChart2.sizeStep)
      }
      return {
        title: this.dataset.chartYAxisTitle2,
        items: arr
      }
    },
    rangeBoundariesChart(){


      let maxBorder = ((obj = this.dataset.charts[0]) => {for (const key in obj) return obj[key];})();
      let minBorder = ((obj = this.dataset.charts[0]) => {for (const key in obj) return obj[key];})();


      maxBorder = this.roundingUp(maxBorder, 1, 2, this.dataset.charts, 1.1);
      minBorder = this.roundingUp(minBorder, -1, 1, this.dataset.charts, 1.2);

      if (!typeof maxBorder!=typeof minBorder){
        minBorder = parseInt(minBorder)
      }
      if(maxBorder.toString().length>minBorder.toString().length+1){
        minBorder = parseInt(minBorder.toString().replace(/^./, '0'))
        minBorder *= 10
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

      let sizeStep =  Number((displayedRange/numberSteps).toFixed(10))


      if(minBorder%sizeStep!=0){
        minBorder = Math.floor(minBorder/sizeStep)*sizeStep
      }

      maxBorder = Number((displayedRange+minBorder).toFixed(10));


      return {
        maxBorder: maxBorder,
        minBorder: minBorder,
        displayedRange: displayedRange,
        numberSteps: numberSteps,
        sizeStep: sizeStep
      }

    },
    rangeBoundariesChart2(){


      let maxBorder = ((obj = this.dataset.charts2[0]) => {for (const key in obj) return obj[key];})();
      let minBorder = ((obj = this.dataset.charts2[0]) => {for (const key in obj) return obj[key];})();


      maxBorder = this.roundingUp(maxBorder, 1, 2, this.dataset.charts2, 1.1);
      minBorder = this.roundingUp(minBorder, -1, 1, this.dataset.charts2, 1.2);

      if (!typeof maxBorder!=typeof minBorder){
        minBorder = parseInt(minBorder)
      }
      if(maxBorder.toString().length>minBorder.toString().length+1){
        minBorder = parseInt(minBorder.toString().replace(/^./, '0'))
        minBorder *= 10
      }


      let displayedRange = this.searchRange(maxBorder, minBorder)
      let numberSteps = this.rangeBoundariesChart.numberSteps
      let x = parseInt(displayedRange.toString().slice(0,2))
      if (x%numberSteps!=0){
        displayedRange=parseInt(x+(numberSteps-x%numberSteps)+displayedRange.toString().slice(2))
      }
      let sizeStep = displayedRange/numberSteps
      if(minBorder%sizeStep!=0){
        minBorder = Math.floor(minBorder/sizeStep)*sizeStep
      }
      maxBorder = displayedRange+minBorder




      return {
        maxBorder: maxBorder,
        minBorder: minBorder,
        displayedRange: displayedRange,
        numberSteps: numberSteps,
        sizeStep: sizeStep
      }

    }
  },
  methods:{
    validationDisplayedRange(value) {
      const simpleNumber = [11, 13, 17, 19, 22, 23, 26, 29, 31, 33, 34, 37, 38, 39, 41, 43, 44, 46, 47, 51, 52, 53, 55, 57, 58,
        59, 61, 62, 65, 66, 67, 68, 69, 71, 73, 74, 76, 77, 78, 79, 82, 83, 85, 86, 87, 88, 89, 91, 92, 93, 94, 95, 97, 99]

      if (simpleNumber.includes(parseInt(value.toString().slice(0, 2)))) return false
      return true
    },
    lastDigit(n) {
      if(n % 1 != 0){
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
    /**
     * bigWay = 1; -1
     * */
    roundingUp(startValue, bigWay, charactersBeforeRounding, arrs, roundingFactor = 1.0){
      let thisValue = startValue
      for (const arr of arrs) {
        for (const idx in arr){
          const value = arr[idx]
          if (thisValue*bigWay<value*bigWay){
            thisValue=value
          }
        }
      }
      if(bigWay==1){
        thisValue = thisValue * roundingFactor
      }
      if(bigWay==-1){
        thisValue = thisValue / roundingFactor
      }
      if (thisValue >= 10){
        let factor
        if (thisValue < 100){
          factor = Math.pow(10, Math.floor(Math.log10(Math.abs(Math.round(thisValue)))));
        }
        else{
          factor = Math.pow(10, Math.floor(Math.log10(Math.abs(Math.round(thisValue)))) - (charactersBeforeRounding-1))
        }
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
  },



}
</script>

<style scoped>
.fourChartPage{
  height: 100%;
  padding: 20px 30px;
  box-sizing: border-box;
}
</style>