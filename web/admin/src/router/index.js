import VueRouter from "vue-router";
import Vue from "vue";
import Admin from "../views/Admin.vue";
import Login from "../views/Login.vue";

// 管理页面组件
import Index from "../components/admin/Index.vue";
import AddArticle from "@/components/article/AddArticle";
import ArticleList from "@/components/article/ArticleList";
import CateList from "@/components/category/CateList";
import UserList from "@/components/user/UserList";
import Comment from "@/components/comment/commentList";
import Profile from '@/components/user/Profile'

// // 路由重复点击捕获错误
// const originalPush = VueRouter.prototype.push
// VueRouter.prototype.push = function push(location, onResolve, onReject) {
//   if (onResolve || onReject) return originalPush.call(this, location, onResolve, onReject)
//   return originalPush.call(this, location).catch(err => err)
// }


Vue.use(VueRouter);

const routes = [
  {
    path: "/login",
    name: "login",
    component: Login,
  },
  {
    path: "/admin",
    name: "admin",
    component: Admin,
    children: [
      { path: "index", component: Index },
      { path: "addArticle", component: AddArticle },
      { path: "addArticle/:id", component: AddArticle, props:true },
      { path: "artList", component: ArticleList },
      { path: "cateList", component: CateList },
      { path: "userList", component: UserList },
      { path: "comment", component: Comment },
      { path: "profile", component: Profile },
    ],
  },
];

const router = new VueRouter({
  routes,
});

router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem("token");
  if (to.name === "login"){
    return next();
  }else{
    if(!token){
      next({ name: "login" });
    }else{
      next();
    }
  } 
});

export default router;
