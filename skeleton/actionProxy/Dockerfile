#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# Dockerfile for docker skeleton (useful for running blackbox binaries, scripts, or Python 3 actions) .
FROM sgx 

USER root
# Upgrade and install basic Python dependencies.
# python3
RUN apt install -y python3 python3-distutils wget
RUN cd /usr/bin && ln -s python3 python
RUN set -ex; wget -O get-pip.py https://github.com/pypa/get-pip/raw/1fe530e9e3d800be94e04f6428460fc4fb94f5a9/get-pip.py; \
  python get-pip.py --disable-pip-version-check "pip==20.1"

RUN apt install -y perl jq zip git curl wget openssl ca-certificates sed openssh-client \
  && update-ca-certificates \
  && pip install --upgrade pip setuptools six \
  && pip install --no-cache-dir gevent==1.3.6 flask==1.0.2 

ENV FLASK_PROXY_PORT 8080

RUN mkdir -p /actionProxy/owplatform
ADD actionproxy.py /actionProxy/
ADD owplatform/__init__.py /actionProxy/owplatform/
ADD owplatform/knative.py /actionProxy/owplatform/
ADD owplatform/openwhisk.py /actionProxy/owplatform/

RUN mkdir -p /action
ADD stub.sh /action/exec
RUN chmod +x /action/exec

CMD ["/bin/bash", "-c", "cd actionProxy && python -u actionproxy.py"]
