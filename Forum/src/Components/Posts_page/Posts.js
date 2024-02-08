// src/App.js
import React, { useState } from "react";
import "./Posts.css";
import axios from "axios";
import { MdOutlineDeleteOutline } from "react-icons/md";

import { FaRegComment } from "react-icons/fa";

import { CiEdit } from "react-icons/ci";
import { useEffect } from "react";

const Posts = (props) => {
  const [showReply, setShow_reply] = useState(true);
  const [show_specific_reply, set_show_specific_reply] = useState(null);
  const [title, setTitle] = useState("");
  const [para, setPara] = useState("");
  const [write_comment, set_write_comment] = useState("");
  const [reload, set_reload] = useState(false);
  const [posts, setPosts] = useState([]);

  const [comments, setComments] = useState([]);

  const [editedPostId, setEditedPostId] = useState(null); // new state variable to store the ID of the post being edited
  useEffect(() => {
    if (posts) {
      console.log("sss ", posts);
      for (var i = 0; i < posts.length; i++) {
        if (posts[i].id === editedPostId) {
          setEditTitle(posts?.[i]?.title);
          setEditContent(posts?.[i]?.content);
        }
      }
    }
  }, [editedPostId]);

  const [editTitle, setEditTitle] = useState(""); // new state variable for editing title
  const [editContent, setEditContent] = useState(""); // new state variable for editing content
  const [editMode, setEditMode] = useState(false); // new state variable to track editing mode

  useEffect(() => {
    // Fetch posts and comments when component mounts
    fetchPosts();
    fetchComments();
  }, [reload]);

  const fetchPosts = async () => {
    try {
      const response = await axios.get("http://localhost:4090/getPost");

      setPosts(response.data);

      //   setPosts(response.data);
    } catch (error) {
      console.error("Error fetching posts:", error);
    }
  };

  const fetchComments = async () => {
    try {
      const response = await axios.get("http://localhost:4090/getComment");
      setComments(response.data);
    } catch (error) {
      console.error("Error fetching comments:", error);
    }
  };

  const Send_post = async () => {
    const userData = {
      Title: title,
      Content: para,
      Email: props.email_on,
    };

    try {
      const response = await axios.post(
        "http://localhost:4090/createPost",
        userData
      );
      console.log("post created ");
      setPara("");
      setTitle("");
      set_reload(!reload);
    } catch (error) {
      console.error("Error fetching comments:", error);
    }
  };

  const show_Comment_button = (id) => {
    if (show_specific_reply === id) {
      set_show_specific_reply(null);
    } else {
      set_show_specific_reply(id);
    }
  };

  const submit_comment = async (post_id) => {
    if (write_comment === "") {
      alert("enter the comment");
    } else {
      const userData = {
        Postid: post_id,
        Content: write_comment,
        Email: props.email_on,
      };

      try {
        const response = await axios.post(
          "http://localhost:4090/createComment",
          userData
        );
        console.log("comment created ");
        set_write_comment("");
        set_reload(!reload);
      } catch (error) {
        console.error("Error fetching comments:", error);
      }
    }
  };

  const submit_delete = async (id) => {
    const userData = {
      Id: id,
    };

    try {
      const response = await axios.post(
        "http://localhost:4090/deletePost",
        userData
      );
      console.log("post deleted ");
      set_reload(!reload);
    } catch (error) {
      console.error("Error fetching comments:", error);
    }
  };

  const submit_edit = async (id) => {
    const userData = {
      Title: editTitle,
      Content: editContent,
      Id: id,
    };

    try {
      const response = await axios.post(
        "http://localhost:4090/editPost",
        userData
      );
      console.log("post edited ");
      setEditedPostId(null);
      setEditMode(false);
      set_reload(!reload);
    } catch (error) {
      console.error("Error fetching comments:", error);
    }
  };

  const emptyState = (e) => {
    setPara("");
    setTitle("");
  };

  const Logout =()=>{
    window.location.reload(false);
  }
  return (
    <div style={{background:"#f8edeb"}}>
      <div style={{display:"flex",justifyContent:"right"}}>
        <button className="logout" onClick={(e)=>Logout()} style={{width:"100px"}}>Logout</button>
      </div>
      <h1 style={{ marginLeft: "30px" }}>Discussion Forum</h1>

      <div className="container">
        <div  className="form-group">
          <label htmlFor="inputText">Title:</label>
          <input
           style={{background:"#f8edeb"}}
            onChange={(e) => setTitle(e.target.value)}
            value={title}
            type="text"
            id="inputText"
            placeholder="Enter text..."
          />
        </div>
        <div className="form-group-para">
          <label htmlFor="inputText">Content:</label>
          <input
           style={{background:"#f8edeb"}}
            onChange={(e) => setPara(e.target.value)}
            value={para}
            type="text"
            id="inputText"
            placeholder="Enter text..."
          />
        </div>
      </div>
      <div className="block-button">
        <button onClick={(e) => Send_post()} className="btn">
          Post
        </button>
        <button onClick={(e) => emptyState(e)} className="btn-c">
          Cancle
        </button>
      </div>

      {posts.map((data, id) => (
        <div key={id} className="data_posts">
          {editMode && editedPostId === data.id ? (
            <>
              <div>
                <input
                  type="text"
                  value={editTitle}
                  onChange={(e) => setEditTitle(e.target.value)}
                />
              </div>
              <div style={{ marginLeft: "32px", fontSize: "12px" }}>
                Posted by : {data.username}
              </div>
              <div>
                {" "}
                <input
                  type="text"
                  value={editContent}
                  onChange={(e) => setEditContent(e.target.value)}
                />
              </div>
              <div>
                <div>
                  <button onClick={() => submit_edit(data.id)} className="btn">
                    Save Changes
                  </button>
                  <button
                    onClick={() => {
                      setEditedPostId(null);
                      setEditMode(false);
                    }}
                    className="btn-c"
                  >
                    Cancel
                  </button>
                </div>
              </div>
            </>
          ) : (
            <>
              {" "}
              <div
                style={{
                  color: "#0056b3",
                  fontSize: "32px",
                  marginLeft: "30px",
                  marginTop: "5px",
                  fontStyle: "normal",
                  fontFamily: "sans-serif",
                }}
              >
                {data.title}
              </div>
              <div style={{ marginLeft: "32px", fontSize: "12px" }}>
                Posted by : {data.username}
              </div>
              <div
                style={{
                  marginLeft: "30px",
                  marginTop: "10px",
                  fontSize: "22px",
                }}
              >
                {data.content}
              </div>
            </>
          )}

          <div style={{ display: "flex", width: "26%", marginTop: "15px" }}>
            <div
              onClick={(e) => show_Comment_button(data.id)}
              style={{ cursor: "pointer", marginLeft: "34px" }}
            >
              <FaRegComment size={30} />
            </div>
            {data.email === props.email_on && (
              <>
                {" "}
                <div
                  onClick={(e) => submit_delete(data.id)}
                  style={{ cursor: "pointer", marginLeft: "34px" }}
                >
                  <MdOutlineDeleteOutline size={30} />
                </div>
                <div
                  onClick={(e) => {
                    setEditedPostId(data.id);
                    setEditMode(!editMode);
                  }}
                  style={{ cursor: "pointer ", marginLeft: "34px" }}
                >
                  <CiEdit size={30} />
                </div>
              </>
            )}
          </div>

          {show_specific_reply === data.id && (
            <div style={{ marginTop: "10px", border: "1px solid gray" }}>
              <div style={{ display: "flex", alignItems: "center" }}>
                <div style={{ marginTop: "10px", width: "400px" }}>
                  <div
                    style={{ marginTop: "12px", marginLeft: "40px" }}
                    className="form-group"
                  >
                    <input
                      onChange={(e) => set_write_comment(e.target.value)}
                      value={write_comment}
                      type="text"
                      id="inputText"
                      placeholder="Enter Comment..."
                    />
                  </div>
                </div>
                <div style={{ marginTop: "2px" }}>
                  <button
                    onClick={(e) => submit_comment(data.id)}
                    className="btn"
                  >
                    Comment
                  </button>
                </div>
              </div>

              <div style={{ marginTop: "7px" }}>
                {comments.map((comment, com_id) => (
                  <>
                    {showReply && comment.postid == data.id && (
                      <div
                        style={{ marginTop: "4px", marginLeft: "40px" }}
                        key={com_id}
                      >
                        <div style={{ fontSize: "26px" }}>
                          {comment.username}
                        </div>
                        <div style={{ fontSize: "17px" }}>
                          {comment.content}
                        </div>
                      </div>
                    )}
                  </>
                ))}
              </div>
            </div>
          )}
        </div>
      ))}
    </div>
  );
};

export default Posts;
