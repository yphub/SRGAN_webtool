<template>
  <div class="ViewerSub" @mousedown="onMouseDown" @mousewheel="onWheel">
    <slot></slot>
    <div class="viewer-container" :style="{transform:`scale(${realImgScale})`}">
      <img
        ref="image"
        :src="src"
        :style="{transform:`translate(${value.imgTx}%,${value.imgTy}%)`}"
        :class="{notrans: value.mousedown}"
        class="viewer-img"
        @load="onImageLoad"
      >
    </div>    
  </div>
</template>

<script>
export default {
  name: "ViewerSub",
  props: {
    value: Object,
    src: String
  },
  data() {
    return {
      imgScaleTimes: 1
    };
  },
  computed: {
    realImgScale() {
      return this.value.imgScale * this.imgScaleTimes;
    }
  },
  methods: {
    onImageLoad() {
      this.$emit("load");

      this.value.imgScale = 0.9;
      this.value.imgTx = 0;
      this.value.imgTy = 0;
      this.value.layer.$emit("layer");
    },
    onMouseDown(e) {
      this.value.mousedown = true;
      window.addEventListener("mouseup", this.onMouseUp);
      window.addEventListener("mousemove", this.onMouseMove);
      //   window.addEventListener("mousewheel", this.onToolWheel);
    },
    onMouseUp(e) {
      this.value.mousedown = false;
      this.replaceImage();
      window.removeEventListener("mouseup", this.onMouseUp);
      window.removeEventListener("mousemove", this.onMouseMove);
      //   window.removeEventListener("mousewheel", this.onToolWheel);
    },
    onMouseMove(e) {
      this.value.moved = true;
      this.value.imgTx +=
        (e.movementX / this.realImgScale / this.$refs.image.clientWidth) * 100;
      this.value.imgTy +=
        (e.movementY / this.realImgScale / this.$refs.image.clientHeight) * 100;
    },
    onWheel(e) {
      this.value.imgScale += e.deltaY / 1000;
      if (this.value.imgScale <= 0.1) {
        this.value.imgScale = 0.1;
      }
      if (!this.value.mousedown) {
        if (this._replaceTimerid != undefined) {
          clearTimeout(this._replaceTimerid);
        }
        this._replaceTimerid = setTimeout(() => {
          this.replaceImage();
          delete this._replaceTimerid;
        }, 250);
      }
    },
    replaceImage() {
      let xInterval =
        (1 -
          this.$el.clientWidth /
            this.$refs.image.clientWidth /
            this.realImgScale) *
        50;
      if (xInterval < 0) {
        this.value.imgTx = 0;
      } else if (this.value.imgTx > xInterval) {
        this.value.imgTx = xInterval;
      } else if (this.value.imgTx < -xInterval) {
        this.value.imgTx = -xInterval;
      }

      let yInterval =
        (1 -
          this.$el.clientHeight /
            this.$refs.image.clientHeight /
            this.realImgScale) *
        50;
      if (yInterval < 0) {
        this.value.imgTy = 0;
      } else if (this.value.imgTy > yInterval) {
        this.value.imgTy = yInterval;
      } else if (this.value.imgTy < -yInterval) {
        this.value.imgTy = -yInterval;
      }
    },
    caculateScaleTimes() {
      let $el = this.$el;
      let $im = this.$refs.image;
      this.imgScaleTimes = Math.min(
        $el.clientHeight / $im.clientHeight,
        $el.clientWidth / $im.clientWidth
      );
    }
  },
  mounted() {
    this.value.layer.$on("layer", this.caculateScaleTimes);
  },
  beforeDestroy() {
    this.value.layer.$off("layer", this.caculateScaleTimes);
  }
};
</script>

<style lang="scss">
.ViewerSub {
  flex-grow: 1;
  overflow: hidden;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  .viewer-container {
    img {
      width: auto;
      height: auto;
      transition: transform ease-in-out 0.4s;
      &.notrans {
        transition: none;
      }
    }
  }
}
</style>