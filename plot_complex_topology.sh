#!/bin/bash
#
# complex_topology.go structure visualiser
#
# how to install graphviz
#
# http://linux.cloudypoint.com/forums/topic/ubuntu_16-04-solved-how-to-install-pydot-and-graphviz/
#
# sudo apt install python-pydot python-pydot-ng graphviz
#
# then create file
# garaph.gv
# with this content
# digraph graphname {
#  a -> b;
# }
#
# and put a command to prompt
# $dot -v -Tpng -ograph graph.gv
# There is a graph.png as a result should appear
#
echo "Creating the diagram of complex_topology.go"
dot -v -Tpng -ocomplex_topology.png complex_topology.gv
echo "Done!"

