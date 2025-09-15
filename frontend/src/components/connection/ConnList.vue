<script setup lang="ts">
import {computed, onMounted, reactive, ref} from "vue";
import {backend} from "../../../wailsjs/go/models.js";
import {idStore} from "../../config/Store.js";
import {AddConn, CreateClient, DeleteConn, Disconnect, QueryConns} from "../../../wailsjs/go/backend/DbService.js";
import {ElMessage} from "element-plus";
import CollapseBtn from "./CollapseBtn.vue";
import Conn = backend.Conn;

const connList = ref<Conn[]>([])

const addFormVisible = ref<boolean>(false);
const conn = ref<Conn>(new Conn());
const connId = ref<number>(0);
//已连接
const linkList = reactive<Array<number>>([])

//菜单折叠
const isCollapse = ref(false);
const asideWidth = computed(() => isCollapse.value ? "60px" : "200px");

async function refreshList() {
  connList.value = await QueryConns()
}

onMounted(() => refreshList())

async function addConn() {
  //入库
  await AddConn(conn.value)
  await refreshList()
  addFormVisible.value = false
  conn.value = new Conn()
}

//列表点击
function connClick(id: number) {
  idStore().connId = id
  connId.value = id
}

//连接按钮文本
const linkBtn = computed(() => linkList.includes(connId.value) ? '断开' : '连接')
const btnType = computed(() => linkList.includes(connId.value) ? 'danger' : 'success')

async function connLink() {
  //已连接，断开
  const id = connId.value;
  try {
    if (linkList.includes(id)) {
      await Disconnect(id)
      ElMessage.success('已断开')
      linkList.splice(linkList.indexOf(id), 1)
      return
    }

    const target = connList.value.find(c => c.Id === id);
    await CreateClient(target);
    ElMessage.success('连接成功')
    linkList.push(id)
  } catch (e: any) {
    ElMessage.error(e.message || e)
  }
}

//右键菜单
async function delConn(id: number) {
  await DeleteConn(id)
  await refreshList()
}
</script>

<template>
  <el-aside :width="asideWidth" class="sidebar">
    <h4>{{ isCollapse ? "连接" : "连接列表" }}</h4>
    <div class="header">
      <el-button type="primary"
                 @click="addFormVisible = true"
                 v-show="!isCollapse"
                 class="btn-add">添加
      </el-button>
      <el-button :type="btnType" @click="connLink" class="btn-link">
        {{ linkBtn }}
      </el-button>
    </div>
    <el-menu :collapse="isCollapse"
             active-text-color="blue"
             default-active="1"
             class="menu">
      <el-menu-item v-for="(conn,idx) in connList"
                    :index="String(idx)"
                    @click="connClick(conn.Id)">
        <el-dropdown trigger="contextmenu">
          <div class="conn-name">{{ conn.Name }}</div>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="delConn(conn.Id)">删除</el-dropdown-item>
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
    <el-form :model="conn">
      <el-form-item label="名称">
        <el-input v-model="conn.Name"/>
      </el-form-item>
      <el-form-item label="地址">
        <el-input v-model="conn.URL"/>
      </el-form-item>
      <el-form-item label="Token">
        <el-input v-model="conn.Token"/>
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="addFormVisible = false">取消</el-button>
        <el-button type="primary" @click="addConn">确定</el-button>
      </div>
    </template>
  </el-dialog>

</template>

<style scoped>
.sidebar {
  height: 100%;
  display: flex;
  flex-direction: column;
  overflow: hidden;
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