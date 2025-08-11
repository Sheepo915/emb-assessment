const showing = document.getElementById("showing");
const sort = document.getElementById("sort");
const goToTop = document.getElementById("goToTop");
const bookDetail = document.getElementById("bookDetail");
const closeDialogBtn = document.getElementById("closeDialog");

const params = new URLSearchParams(window.location.search);

if (params.has("id")) {
  bookDetail.showModal();
}

if (params.has("limit")) {
  showing.value = params.get("limit");
}

if (params.has("sort")) {
  sort.value = params.get("sort");
}

const updateQuery = () => {
  const params = new URLSearchParams(window.location.search);
  params.set("limit", showing.value);
  params.set("sort", sort.value);
  params.set("page", 1);
  window.location.search = params.toString();
};

showing.onchange = updateQuery;
sort.onchange = updateQuery;

const toTop = () => {
  window.scrollTo({
    top: 0,
    behavior: "smooth",
  });
};

const showGoToTopBtn = () => {
  if (window.scrollY > 10) {
    goToTop.removeAttribute("hidden");
  } else {
    goToTop.setAttribute("hidden", "");
  }
};

window.onscroll = showGoToTopBtn;
goToTop.onclick = toTop;

const openBookDialog = () => {
  bookDetail.showModal();
};

closeDialogBtn.addEventListener("click", () => {
  const params = new URLSearchParams(window.location.search);
  params.delete("id");
  bookDetail.close();
  window.location.search = params.toString();
});

bookDetail.addEventListener("click", (e) => {
  const rect = bookDetail.getBoundingClientRect();
  const isInDialog =
    rect.top <= e.clientY &&
    e.clientY <= rect.top + rect.height &&
    rect.left <= e.clientX &&
    e.clientX <= rect.left + rect.width;
  if (!isInDialog) {
    const params = new URLSearchParams(window.location.search);
    params.delete("id");
    bookDetail.close();
    window.location.search = params.toString();
  }
});

const books = document.querySelectorAll(".book-list .card");

books.forEach((book) => {
  book.onclick = () => {
    const params = new URLSearchParams(window.location.search);
    params.set("id", book.dataset.key);
    window.location.search = params.toString();
  };
});
