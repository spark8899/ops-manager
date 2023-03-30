<template>
  <div class="bottom-info">
    <div>
      <span>Powered by</span>
      <span>
        <a href="https://github.com/spark8899/ops-manager">{{ $OPS_MANAGER.appName }}</a>
      </span>
      <el-divider direction="vertical" />
      <span>Copyright</span>
      <span>
        <a href="https://github.com/spark8899">spark8899</a>
      </span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BottomInfo'
}
</script>

<style lang="scss">
.bottom-info {
  color: #888;
  height: 30px;
  line-height: 12px;
  a {
    color: #888;
  }
  div {
    display: flex;
    justify-content: center;
    span{
      margin: 0 3px;
    }
  }
}
</style>
