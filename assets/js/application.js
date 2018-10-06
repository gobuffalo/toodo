require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

window.watchCheckboxes = function() {
  $(".item-form input[type=checkbox]").on("change", (e) => {
    let $e = $(e.target);
    let $form = $e.closest("form");
    $form.submit();
  });
};

$(() => {
  watchCheckboxes();
});
