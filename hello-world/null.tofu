
resource "time_sleep" "wait_30_seconds" {

  create_duration = "999s"
}
resource "null_resource" "long_operation" {
  provisioner "local-exec" {
    command = "sleep 1200"  // Sleeps for 20 minutes
  }
}

resource "time_sleep" "wait_30_seconds" {
  create_duration = "999s"
}
