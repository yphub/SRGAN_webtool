<template>
  <div
    class="ImgViewer"
    @dragstart="onDragStart"
    @click="onClick"
  >
    <transition-group
      name="viewer"
      tag="div"
      class="viewer-trans"
    >
      <ViewerSub
        v-for="(srceach,index) in src"
        :key="srceach || index"
        :src="srceach"
        v-show="srceach != null"
        :value="Scalevalue"
        :style="{'background-color':background[index].color}"
      >
        <span
          class="lutext"
          v-if="srceach"
        >{{ background[index].lutext }}</span>
        <span class="bgtext">{{ background[index].bgtext }}</span>
      </ViewerSub>
    </transition-group>
  </div>
</template>

<script>
import Vue from "vue";
import ViewerSub from "./ViewerSub.vue";
export default {
  name: "ImgViewer",
  components: {
    ViewerSub
  },
  props: {
    src: Array,
    background: Array
  },
  data() {
    return {
      Scalevalue: {
        moved: false,
        imgScale: 1,
        imgTx: 0,
        imgTy: 0,
        mousedown: false,
        layer: new Vue()
      }
    };
  },
  methods: {
    onClick() {
      if (!this.Scalevalue.moved) {
        if (this._ToolSingleclicked) {
          this.$emit("dbclick");
        } else {
          this._ToolSingleclicked = true;
          setTimeout(() => {
            delete this._ToolSingleclicked;
          }, 300);
        }
      }
      this.Scalevalue.moved = false;
    },
    onDragStart(e) {
      e.preventDefault();
    }
  }
};
</script>

<style lang="scss">
.ImgViewer {
  flex-grow: 1;
  width: 100%;
  // align-items: center;
  .viewer-trans {
    width: 100%;
    height: 100%;
    overflow: hidden;
    display: flex;
    .ViewerSub {
      // border-left: 1px solid gray;
      color: white;
      .bgtext {        
        font-size: 20px;
        font-weight: lighter;
        position: absolute;
        transform: translate(-50%, -50%);
        left: 50%;
        top: 50%;
      }
      .lutext {
        position: absolute;
        left: 0;
        top: 0;
        background-color: rgba(0, 0, 0, 0.5);        
        font-size: 10px;
        font-weight: lighter;
        padding: 2px 5px;
        z-index: 3;
      }
    }
    .ViewerSub:nth-child(1) {
      // border-left: none;
    }
  }
}
</style>
