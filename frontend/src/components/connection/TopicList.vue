<script setup lang="ts">
import {computed, reactive, ref, watchEffect} from "vue";
import {backend} from "../../../wailsjs/go/models.js";
import {idStore} from "../../config/Store.js";
import {
  AddTopic,
  CreateConsumer,
  CreateProducer,
  DelTopic,
  QueryTopics,
  Receive,
  RemoveConsumer,
  RemoveProducer
} from "../../../wailsjs/go/backend/DbService.js";
import {storeToRefs} from "pinia";
import {ElMessage} from "element-plus";
import CollapseBtn from "./CollapseBtn.vue";
import Topic = backend.Topic;

const store = idStore()
const {connId, topicId} = storeToRefs(store)
const topicList = ref<Topic[]>([])
const addFormVisible = ref<boolean>(false);
const topic = ref<Topic>(new Topic());
// 折叠状态
const isCollapse = ref(false);
const asideWidth = computed(() => isCollapse.value ? "60px" : "220px");
//已连接
const linkList = reactive<Array<number>>([])

async function refreshList() {
  topicList.value = await QueryTopics(connId.value)
}

watchEffect(async () => refreshList())

//连接按钮文本
const linkBtn = computed(() => linkList.includes(topicId.value) ? '断开' : '监听')
const btnType = computed(() => linkList.includes(topicId.value) ? 'danger' : 'success')

async function addTopic() {
  topic.value.ConnID = connId.value
  await AddTopic(topic.value)
  //刷新
  await refreshList()
  //关闭弹窗
  addFormVisible.value = false
  //还原topic
  topic.value = new Topic()
}

function topicClick(id: number) {
  idStore().topicId = id
}

async function topicLink() {
  const id = topicId.value
  console.log(linkList.includes(topicId.value))
  try {
    //已连接
    if (linkList.includes(id)) {
      await RemoveConsumer(id)
      await RemoveProducer(id)
      linkList.splice(linkList.indexOf(id), 1)
      ElMessage.success('已断开')
      return
    }
    const topic = topicList.value.find(t => t.Id === id);
    await CreateProducer(topic)
    await CreateConsumer(topic)
    await Receive(id)
    linkList.push(id)
    console.log(linkList.toString())
    ElMessage.success('已开始监听')
  } catch (err: any) {
    ElMessage.error(err)
    console.log(err)
  }
}

//右键菜单
async function delConn(id: number) {
  await DelTopic(id)
  await refreshList()
}
</script>

<template>
    <el-aside :width="asideWidth" class="sidebar">
      <h4>{{ isCollapse ? "Topic" : "Topic列表" }}</h4>
      <div class="header">
        <el-button type="primary"
                   @click="addFormVisible = true"
                   v-if="!isCollapse"
                   class="btn-add">添加
        </el-button>
        <el-button :type="btnType" @click="topicLink"
                   class="btn-link">{{ linkBtn }}
        </el-button>
      </div>

      <el-menu :collapse="isCollapse" class="menu">
        <el-menu-item v-for="topic in topicList"
                      :index="String(topic.Id)"
                      @click="topicClick(topic.Id)">
          <el-dropdown trigger="contextmenu">
            {{ topic.Name }}
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="delConn(topic.Id)">删除</el-dropdown-item>
                <el-dropdown-item @click="">编辑</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </el-menu-item>
      </el-menu>

      <CollapseBtn :isCollapse="isCollapse" @collapse="isCollapse =!isCollapse"></CollapseBtn>
    </el-aside>

  <!--  添加弹窗-->
  <el-dialog v-model="addFormVisible" title="添加连接" width="500px">
    <el-form :model="topic">
      <el-form-item label="名称">
        <el-input v-model="topic.Name"/>
      </el-form-item>
      <el-form-item label="Topic">
        <el-input v-model="topic.Topic"/>
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="addFormVisible = false">取消</el-button>
        <el-button type="primary" @click="addTopic">确定</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<style scoped>
.sidebar {
  height: 100%;
  display: flex;
  transition: width 0.3s;
  flex-direction: column;
  overflow: hidden;
  border-right: #3d525e solid 1px;
  border-left: #3d525e solid 1px;
}

.menu {
  border-right: none;
  width: 100%;
  flex: 1;
}

h4 {
  text-align: center;
  margin: 10px 0;
}

.header {
  position: relative;
  height: 40px;
  padding: 0 10px;
  display: flex;
  align-items: center;
}

.btn-add {
  position: absolute;
  left: 0;
}

.btn-link {
  position: absolute;
  right: 0;
}

.el-button {
  height: 30px;
  min-width: 60px;
}
</style>
