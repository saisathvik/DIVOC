import 'package:json_annotation/json_annotation.dart';

part 'parser.g.dart';

class NavigationFlowParser {
/*  FlowTree buildFlowTree(Map<String, dynamic> jsonMap) {
    FlowTree flowTree;
    jsonMap.forEach((key, value) {
      final metadata = FlowMeta.fromJson(value["metadata"]);
      final nextRoute = value["nextRoutes"];
      if (nextRoute == null) {
        flowTree = FlowTree(metadata, []);
      } else {
        //TODO: nextRoute not populated
        final tree = buildFlowTree(nextRoute);
        flowTree = FlowTree(metadata, tree.nextRoutes);
      }
    });
    return flowTree;
  }*/

  FlowTree buildFlowTree(Map<String, dynamic> jsonMap) {
    return FlowTree.fromJson(jsonMap);
  }
}

@JsonSerializable()
class FlowTree {
  @JsonKey(name: 'metadata')
  final FlowMeta flowMeta;

  @JsonKey(name: 'routeKey')
  final String routeKey;

  @JsonKey(name: 'nextRoutes')
  final List<FlowTree> nextRoutes;

  FlowTree(this.flowMeta, this.routeKey, this.nextRoutes);

  factory FlowTree.fromJson(Map<String, dynamic> json) =>
      _$FlowTreeFromJson(json);

  Map<String, dynamic> toJson() => _$FlowTreeToJson(this);

  FlowTree findNode(FlowTree node, String value) {
    if (node.routeKey == value) {
      return node;
    } else {
      for(FlowTree element in node.nextRoutes){
        final flowTree = findNode(element, value);
        if (flowTree != null) {
          return flowTree;
        }
      }
     /*
      node.nextRoutes.forEach((element) {
        final flowTree = findNode(element, value);
        if (flowTree != null) {
          break;
          return flowTree;
        }
      });*/
    }
    return null;
  }
}

@JsonSerializable()
class FlowMeta {
  final String label;
  final String formKey;

  FlowMeta({this.label, this.formKey});

  factory FlowMeta.fromJson(Map<String, dynamic> json) =>
      _$FlowMetaFromJson(json);

  Map<String, dynamic> toJson() => _$FlowMetaToJson(this);
}